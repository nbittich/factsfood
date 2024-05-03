package services

import (
	"context"

	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/types"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

const UserCollection = "user"

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUser(ctx context.Context, newUserForm *types.NewUserForm) (*types.User, error) {
	collection := db.GetCollection(UserCollection)

	err := ValidateStruct(newUserForm)
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

	// todo
	exist, err := db.Exist(ctx, filter, collection)
	if err != nil {
		return nil, err
	}

	if exist {
		m := types.InvalidMessage{"general": []string{"user.exist"}}
		return nil, types.InvalidFormError{Messages: m}
	}

	user := &types.User{
		Username: newUserForm.Username,
		Password: password,
		Email:    newUserForm.Email,
		Enabled:  false, // FIXME should send an email
	}

	_, err = db.InsertOrUpdate(ctx, user, collection)
	return user, err
}
