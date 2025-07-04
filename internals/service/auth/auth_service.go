package auth_service

import (
	"errors"

	"github.com/aaryansinhaa/patient-management-system/internals/model"
	"github.com/aaryansinhaa/patient-management-system/internals/repositories"
	"github.com/aaryansinhaa/patient-management-system/internals/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo       repositories.UserRepository
	jwtManager *utils.JWTManager
}

func NewAuthService(repo repositories.UserRepository, jwtManager *utils.JWTManager) *authService {
	return &authService{repo: repo, jwtManager: jwtManager}
}

func (s *authService) Register(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.ID = uuid.New()
	return s.repo.CreateUser(*user)
}

func (s *authService) Login(username, password string) (*model.User, string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil || user == nil {
		return nil, "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid username or password")
	}

	token, err := s.jwtManager.Generate(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
