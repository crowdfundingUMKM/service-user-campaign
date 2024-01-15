package core

import (
	"errors"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUsers() ([]User, error)

	GetUserByUnixID(UnixID string) (User, error)
	DeactivateAccountUser(input DeactiveUserInput, adminId string) (bool, error)
	ActivateAccountUser(input DeactiveUserInput, adminId string) (bool, error)

	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	SaveToken(UnixID string, Token string) (User, error)

	// notif
	ReportAdmin(UnixID string, input ReportToAdminInput) (NotifCampaign, error)
	GetAllReports() ([]NotifCampaign, error)

	UpdatePasswordByUnixID(UnixID string, input UpdatePasswordInput) (User, error)
	UpdateUserByUnixID(UnixID string, input UpdateUserInput) (User, error)

	SaveAvatar(UnixID string, fileLocation string) (User, error)

	DeleteToken(UnixID string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserByUnixID(UnixID string) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	if user.UnixID == "" {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}
func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.GetAllUser()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *service) DeactivateAccountUser(input DeactiveUserInput, adminId string) (bool, error) {
	// fin user by unix id
	user, err := s.repository.FindByUnixID(input.UnixID)
	if err != nil {
		return false, err
	}
	if adminId == "" {
		return false, errors.New("Admin ID is empty")
	}
	user.UpdateIdAdmin = adminId
	user.StatusAccount = "deactive"
	_, err = s.repository.UpdateStatusAccount(user)

	if err != nil {
		return false, err
	}

	if user.UnixID == "" {
		return true, nil
	}
	return true, nil
}

func (s *service) ActivateAccountUser(input DeactiveUserInput, adminId string) (bool, error) {
	// fin user by unix id
	user, err := s.repository.FindByUnixID(input.UnixID)
	if err != nil {
		return false, err
	}
	if adminId == "" {
		return false, errors.New("Admin ID is empty")
	}
	user.UpdateIdAdmin = adminId
	user.StatusAccount = "active"
	_, err = s.repository.UpdateStatusAccount(user)

	if err != nil {
		return false, err
	}

	if user.UnixID == "" {
		return true, nil
	}
	return true, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.UnixID = uuid.New().String()[:12]
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.BioUser = input.BioUser
	user.AvatarFileName = "/crwdstorage/avatar_investor/dafault-avatar.png"
	user.FBLink = "https://www.facebook.com/"
	user.IGLink = "https://www.instagram.com/"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	// convert data os env to string
	user.StatusAccount = string(os.Getenv("STATUS_ACCOUNT"))

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

// save token to database
func (s *service) SaveToken(UnixID string, Token string) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	user.Token = Token
	_, err = s.repository.UpdateToken(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Notif
func (s *service) ReportAdmin(UnixID string, input ReportToAdminInput) (NotifCampaign, error) {
	notif := NotifCampaign{}
	notif.UserCampaignId = UnixID
	notif.Title = input.Title
	notif.Description = input.Description
	notif.TypeError = input.TypeError
	notif.StatusNotif = 1

	newNotif, err := s.repository.SaveReport(notif)
	if err != nil {
		return newNotif, err
	}
	return newNotif, nil
}

// get all users
func (s *service) GetAllReports() ([]NotifCampaign, error) {
	report, err := s.repository.GetAllReport()
	if err != nil {
		return report, err
	}
	return report, nil
}

func (s *service) UpdateUserByUnixID(UnixID string, input UpdateUserInput) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	if user.UnixID == "" {
		return user, errors.New("No user found on with that ID")
	}

	user.Name = input.Name
	user.Phone = input.Phone
	user.BioUser = input.BioUser
	user.Address = input.Address
	user.Country = input.Country
	user.FBLink = input.FBLink
	user.IGLink = input.IGLink

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) UpdatePasswordByUnixID(UnixID string, input UpdatePasswordInput) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	if user.UnixID == "" {
		return user, errors.New("No user found on with that ID")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.OldPassword))

	if err != nil {
		return user, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	updatedUser, err := s.repository.UpdatePassword(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) SaveAvatar(UnixID string, fileLocation string) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.UploadAvatarImage(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) DeleteToken(UnixID string) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	if user.UnixID == "" {
		return user, errors.New("No user found on with that ID")
	}

	user.Token = ""

	updatedUser, err := s.repository.UpdateToken(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
