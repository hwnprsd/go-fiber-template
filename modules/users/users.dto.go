package users

type OneUserDto struct {
	Id string `json:"id" validate:"required"`
}
