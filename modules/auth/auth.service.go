package auth

import (
	"flaq.club/backend/models"
	"flaq.club/backend/modules/users"
	. "flaq.club/backend/utils"
)

func Login(data RequestBody) interface{} {
	body := data.Data.(*LoginDto) // Cast to the actual type
	return body.Username
}

func Signup(data RequestBody) interface{} {
	body := data.Data.(*SignupDto)
	// Do JWT crap

	return users.CreateUser(models.User{
		Username:  body.Username,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	})

}
