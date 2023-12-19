package core

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	BioUser  string `json:"bio_user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
