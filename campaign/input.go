package campaign

type RegisterUserInput struct {
	Name        string `json:"name" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Description string `json:"description"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
}
