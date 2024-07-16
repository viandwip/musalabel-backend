package models

import (
	"mime/multipart"
	"time"
)

type User struct {
	Id           string                `json:"id" form:"id"`
	Email        string                `json:"email" form:"email" valid:"required, email"`
	Password     string                `json:"password" form:"password" valid:"required, stringlength(6|100)~Password min. 6 chars"`
	Phone_number string                `json:"phone_number" form:"phone_number" valid:"required"`
	Role         string                `json:"role" form:"role"`
	ImageUpload  *multipart.FileHeader `form:"image"`
	Image        string                `json:"image"`
	Address      string                `json:"address" form:"address"`
	Display_name string                `json:"display_name" form:"display_name"`
	First_name   string                `json:"first_name" form:"first_name"`
	Last_name    string                `json:"last_name" form:"last_name"`
	Birthday     string                `json:"birthday" form:"birthday"`
	Gender       string                `json:"gender" form:"gender"`

	Created_at *time.Time `json:"created_at,omitempty" form:"created_at,omitempty"`
	Updated_at *time.Time `json:"updated_at,omitempty" form:"updated_at,omitempty"`
}
