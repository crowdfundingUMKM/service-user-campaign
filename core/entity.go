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
	ID             int       `json:"id"`
	UnixID         string    `json:"unix_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Country        string    `json:"country"`
	Address        string    `json:"address"`
	BioUser        string    `json:"bio_user"`
	FBLink         string    `json:"fb_link"`
	IGLink         string    `json:"ig_link"`
	PasswordHash   string    `json:"password_hash"`
	StatusAccount  string    `json:"status_account"`
	AvatarFileName string    `json:"avatar_file_name"`
	Token          string    `json:"token"`
	UpdateIdAdmin  string    `json:"update_id_admin"`
	UpdateAtAdmin  time.Time `json:"update_at_admin"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type NotifCampaign struct {
	ID             int       `json:"id"`
	UserCampaignId string    `json:"user_campaign_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TypeError      string    `json:"type_error"`
	Document       string    `json:"document"`
	StatusNotif    int       `json:"status_notif"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
