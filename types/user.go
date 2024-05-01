package types

type NewUserForm struct {
	Username        string `form:"username" json:"username" validate:"required,min=3,max=15"`
	Password        string `form:"password" json:"password" validate:"required,min=6,max=18,password"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" validate:"eqcsfield=Password"`
	Email           string `form:"email" json:"email" validate:"required,email"`
	ConfirmEmail    string `form:"confirmEmail" json:"confirmEmail" validate:"eqcsfield=Email"`
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
