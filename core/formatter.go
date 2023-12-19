package core

type UserCampaignFormatter struct {
	ID            int    `json:"id"`
	UnixID        string `json:"unix_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	BioUser       string `json:"bio_user"`
	Token         string `json:"token"`
	StatusAccount string `json:"status_account"`
}

func FormatterUser(user User, token string) UserCampaignFormatter {
	formatter := UserCampaignFormatter{
		ID:            user.ID,
		UnixID:        user.UnixID,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		BioUser:       user.BioUser,
		Token:         token,
		StatusAccount: user.StatusAccount,
	}
	return formatter
}
