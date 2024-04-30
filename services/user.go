package services

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/types"
)

const UserCollection = "user"

func NewUser(ctx context.Context, newUserForm *types.NewUserForm) (*types.User, error) {
	err := Validate.Struct(newUserForm)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// FIXME should be logged instead
			fmt.Println(err)
			return nil, err
		}

		// FIXME just for testing
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return nil, err
	}
	user := &types.User{
		Username: newUserForm.Username,
		Password: newUserForm.Password,
		Email:    newUserForm.Email,
		Enabled:  false, // FIXME should send an email
	}

	_, err = db.InsertOrUpdate(ctx, user, UserCollection)
	return user, err
}
