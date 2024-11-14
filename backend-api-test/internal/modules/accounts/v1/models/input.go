package models

type AuthLogin struct {
	Email    *string `json:"email" binding:"required"`
	Password *string `json:"password" binding:"required,min=5,max=254"`
}
