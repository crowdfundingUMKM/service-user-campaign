package core

type UserCampaignFormatter struct {
	ID            int    `json:"id"`
	UnixID        string `json:"unix_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Description   string `json:"description"`
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
		Description:   user.Description,
		Token:         token,
		StatusAccount: user.StatusAccount,
	}
	return formatter
}
