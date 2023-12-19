package core

import "time"

// type User struct {
// 	ID             int
// 	UnixID         string
// 	Name           string
// 	Email          string
// 	Phone          string
// 	Description    string
// 	PasswordHash   string
// 	AvatarFileName string
// 	StatusAccount  string
// 	Token          string
// 	CreatedAt      time.Time
// 	UpdatedAt      time.Time
// }

type User struct {
	ID             int
	UnixID         string
	Name           string
	Email          string
	Phone          string
	Country        string
	Address        string
	BioUser        string
	FBLink         string
	IGLink         string
	PasswordHash   string
	StatusAccount  string
	AvatarFileName string
	Token          string
	UpdateIDAdmin  string
	UpdateAtAdmin  time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type NotifCampaign struct {
	ID           int
	UserInvestor string
	Title        string
	Description  string
	TypeError    string
	Document     string
	StatusNotif  int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
