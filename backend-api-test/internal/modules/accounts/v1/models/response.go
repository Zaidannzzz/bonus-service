package models

type UserResponse struct {
	UUID      *string     `json:"uuid"`
	Email     *string     `json:"email"`
	FullName  *string     `json:"full_name"`
	LastLogin *uint       `json:"last_login"`
	Gender    *GenderType `json:"gender"`
	CreatedAt *uint       `json:"created_at"`
	UpdatedAt *uint       `json:"updated_at"`
}

type UserProfileResponse struct {
	UUID     *string `json:"uuid"`
	Email    *string `json:"email"`
	FullName *string `json:"full_name"`
	Photo    *string `json:"photo"`
}

type Cookies struct {
	AccessToken string `json:"access_token"`
}
