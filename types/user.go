package types

type NewUserForm struct {
	Username        string `json:"username" form:"username" validate:"required,min=3,max=15"`
	Password        string `json:"password" form:"password" validate:"required,min=6,max=18,password"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"eqcsfield=Password"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	ConfirmEmail    string `json:"confirmEmail" form:"confirmEmail" validate:"eqcsfield=Email"`
}

type User struct {
	ID       string      `bson:"_id" json:"_id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Enabled  bool        `json:"enabled"`
	Email    string      `json:"email"`
	Profile  UserProfile `json:"profile"`
	Settings UserSetting `json:"settings"`
}

type UserProfile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"LastName"`
}

type UserSetting struct {
	Lang string `json:"lang"`
}

func (user *User) GetID() string {
	return user.ID
}

func (user *User) SetID(id string) {
	user.ID = id
}
