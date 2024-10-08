package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/services/email"
	"github.com/nbittich/factsfood/services/utils"
	"github.com/nbittich/factsfood/types"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserCollection              = "user"
	UserActivationURLCollection = "userActivationUrl"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUser(ctx context.Context, newUserForm *types.NewUserForm) (*types.User, error) {
	lang := ctx.Value(types.LangKey).(string)
	collection := db.GetCollection(UserCollection)
	err := utils.ValidateStruct(newUserForm)
	if err != nil {
		return nil, err
	}

	password, err := hashPassword(newUserForm.Password)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"$or": []bson.M{
			{"email": newUserForm.Email},
			{"username": newUserForm.Username},
		},
	}

	exist, err := db.Exist(ctx, filter, collection)
	if err != nil {
		return nil, err
	}

	if exist {
		m := types.InvalidMessage{"general": "home.signup.user.exist"}
		return nil, types.InvalidFormError{Form: newUserForm, Messages: m}
	}

	user := &types.User{
		Username: newUserForm.Username,
		Password: password,
		Email:    newUserForm.Email,
		Enabled:  false,
		Settings: types.UserSetting{Lang: lang},
		Profile:  types.UserProfile{},
		Roles:    []types.Role{types.USER},
	}

	go sendActivationEmail(user, true)
	return user, nil
}

func sendActivationEmail(user *types.User, createUser bool) {
	collection := db.GetCollection(UserCollection)

	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	if createUser {
		_, e := db.InsertOrUpdate(ctx, user, collection)
		if e != nil {
			log.Println("could not create user:", e)
			return
		}
	}
	activateURL, e := GenerateActivateURL(ctx, config.BaseURL+"/users/activate", user.ID)
	if e != nil {
		log.Println("error while generating validation url", e)
		return
	}
	email.SendAsync([]string{user.Email}, []string{}, "Activate your account", fmt.Sprintf(`<a href="%s">Activate your account now!</p>`, activateURL))
}

func FindByUsernameOrEmail(ctx context.Context, username string) (types.User, error) {
	userCollection := db.GetCollection(UserCollection)
	filter := bson.M{
		"$or": []bson.M{
			{"email": username},
			{"username": username},
		},
	}
	return db.FindOneBy[types.User](ctx, filter, userCollection)
}

func ActivateUser(ctx context.Context, hash string) (bool, error) {
	userCollection := db.GetCollection(UserCollection)
	userActivationURLCollection := db.GetCollection(UserActivationURLCollection)
	userActivationURL, err := db.FindOneBy[types.UserActivationURL](ctx, bson.M{
		"hash": hash,
	}, userActivationURLCollection)
	if err != nil {
		return false, err
	}
	user, err := db.FindOneByID[types.User](ctx, userCollection, userActivationURL.UserID)
	if err != nil {
		return false, err
	}
	if user.Enabled {
		return false, fmt.Errorf("user already enabled")
	}
	now := time.Now()
	duration := now.Sub(userActivationURL.UpdatedAt)
	if duration > config.ActivationExpiration {
		log.Println("activation link no longer valid")
		go sendActivationEmail(&user, false)
		return false, fmt.Errorf("invalid hash")
	}
	userActivationURL.UpdatedAt = now

	user.Enabled = true
	_, err = db.InsertOrUpdate(ctx, &user, userCollection)
	if err != nil {
		return false, err
	}
	_, _ = db.InsertOrUpdate(ctx, &userActivationURL, userActivationURLCollection)
	return true, nil
}

func GenerateActivateURL(ctx context.Context, baseURL string, userID string) (string, error) {
	userCollection := db.GetCollection(UserCollection)
	userActivationURLCollection := db.GetCollection(UserActivationURLCollection)
	user, err := db.FindOneByID[types.User](ctx, userCollection, userID)
	if err != nil {
		return "", err
	}
	if user.Enabled {
		return "", fmt.Errorf("user.alreadyEnabled")
	}
	filter := bson.M{
		"userId": user.ID,
	}
	userActivationURL, err := db.FindOneBy[types.UserActivationURL](ctx, filter, userActivationURLCollection)
	if err != nil {
		now := time.Now()
		duration := now.Sub(userActivationURL.UpdatedAt)
		if duration < config.ActivationExpiration {
			log.Println("activation link still valid")
			return userActivationURL.GenerateURL(baseURL), nil
		}
	}
	userActivationURL.Hash = uuid.New().String()
	userActivationURL.UpdatedAt = time.Now()
	userActivationURL.UserID = userID
	_, err = db.InsertOrUpdate(ctx, &userActivationURL, userActivationURLCollection)
	if err != nil {
		return "", nil
	}
	return userActivationURL.GenerateURL(baseURL), nil
}
