package dto

type UserDto struct {
	UUID     string `json:"uuid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
