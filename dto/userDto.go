package dto

type UserDto struct {
	UUID     string `json:"uuid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserInfoDto struct {
	UserName  string `json:"username"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUserPreferences struct {
	Preference string `json:"preference"`
	UpdatedAt  string `json:"updated_at"`
}
