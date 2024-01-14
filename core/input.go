package core

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	BioUser  string `json:"bio_user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckPhoneInput struct {
	Phone string `json:"phone" binding:"required"`
}

type DeactiveUserInput struct {
	UnixID string `json:"unix_id" binding:"required"`
}

type ActiveUserInput struct {
	UnixID string `json:"unix_id" binding:"required"`
}

//	type DeleteUserInput struct {
//		UnixID string `json:"unix_id" binding:"required"`
//	}
//
// Not user binding:"required" because data can null
type UpdateUserInput struct {
	Name    string `json:"name" `
	Phone   string `json:"phone" `
	BioUser string `json:"bio_user" `
	Addreas string `json:"addreas" `
	Country string `json:"country" `
	FBLink  string `json:"fb_link" `
	IGLink  string `json:"ig_link" `
}

type UpdatePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ReportToAdminInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	TypeError   string `json:"type_error" binding:"required"`
}

type GetUserIdInput struct {
	UnixID string `uri:"unix_id" binding:"required"`
}
