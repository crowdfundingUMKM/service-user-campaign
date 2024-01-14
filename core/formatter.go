package core

type UserCampaignFormatter struct {
	ID             int    `json:"id"`
	UnixID         string `json:"unix_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	BioUser        string `json:"bio_user"`
	Token          string `json:"token"`
	StatusAccount  string `json:"status_account"`
	AvatarFileName string `json:"avatar_file_name"`
}

func FormatterUser(user User, token string) UserCampaignFormatter {
	formatter := UserCampaignFormatter{
		ID:             user.ID,
		UnixID:         user.UnixID,
		Name:           user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		BioUser:        user.BioUser,
		Token:          token,
		StatusAccount:  user.StatusAccount,
		AvatarFileName: user.AvatarFileName,
	}
	return formatter
}

type UserDetailFormatter struct {
	ID             int    `json:"id"`
	UnixID         string `json:"unix_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	BioUser        string `json:"bio_user"`
	Address        string `json:"address"`
	Country        string `json:"country"`
	FBLink         string `json:"fb_link"`
	IGLink         string `json:"ig_link"`
	Token          string `json:"token"`
	StatusAccount  string `json:"status_account"`
	AvatarFileName string `json:"avatar_file_name"`
}

func FormatterUserDetail(user User, updatedUser User) UserDetailFormatter {
	formatter := UserDetailFormatter{
		ID:             user.ID,
		UnixID:         user.UnixID,
		Name:           user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		BioUser:        user.BioUser,
		Address:        user.Address,
		Country:        user.Country,
		FBLink:         user.FBLink,
		IGLink:         user.IGLink,
		StatusAccount:  user.StatusAccount,
		AvatarFileName: user.AvatarFileName,
	}
	// read data before update if null use old data
	if updatedUser.Name != "" {
		formatter.Name = updatedUser.Name
	}
	if updatedUser.Phone != "" {
		formatter.Phone = updatedUser.Phone
	}
	if updatedUser.BioUser != "" {
		formatter.BioUser = updatedUser.BioUser
	}
	if updatedUser.AvatarFileName != "" {
		formatter.AvatarFileName = updatedUser.AvatarFileName
	}
	if updatedUser.StatusAccount != "" {
		formatter.StatusAccount = updatedUser.StatusAccount
	}
	if updatedUser.Address != "" {
		formatter.Address = updatedUser.Address
	}
	if updatedUser.Country != "" {
		formatter.Country = updatedUser.Country
	}
	if updatedUser.FBLink != "" {
		formatter.FBLink = updatedUser.FBLink
	}
	if updatedUser.IGLink != "" {
		formatter.IGLink = updatedUser.IGLink
	}
	return formatter
}

// Notif
// notify formater
type NotifyFormatter struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TypeError   string `json:"type_error"`
	StatusNotif int    `json:"status_notif"`
}

func FormatterNotify(notify NotifCampaign) NotifyFormatter {
	formatter := NotifyFormatter{
		ID:          notify.ID,
		Title:       notify.Title,
		Description: notify.Description,
		TypeError:   notify.TypeError,
		StatusNotif: notify.StatusNotif,
	}
	return formatter
}

// get user admin status
type UserCampaign struct {
	UnixCampaign       string `json:"unix_admin"`
	StatusAccountAdmin string `json:"status_account_admin"`
}

// get user campaign status
func FormatterUserCampaignID(user User) UserCampaign {
	formatter := UserCampaign{
		UnixCampaign:       user.UnixID,
		StatusAccountAdmin: user.StatusAccount,
	}
	return formatter
}
