package campaign

import "time"

type User struct {
	ID             int
	UnixID         string
	Name           string
	Email          string
	Phone          string
	Description    string
	PasswordHash   string
	AvatarFileName string
	StatusAccount  string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
