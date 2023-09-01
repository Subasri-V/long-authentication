package models

type UserInput struct {
	Name            string `json:"name" bson:"name"`
	Email           string `json:"email" bson:"email"`
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirmpassword" bson:"confirmpassword"`
}
