package services

import (
	"github.com/LamichhaneBibek/go-template/internal/api/models"
	"github.com/LamichhaneBibek/go-template/internal/config"
	"github.com/LamichhaneBibek/go-template/internal/constants"
	"github.com/LamichhaneBibek/go-template/internal/db"
	"github.com/LamichhaneBibek/go-template/internal/utils"
	"github.com/LamichhaneBibek/go-template/internal/utils/logger"
	"github.com/LamichhaneBibek/go-template/internal/views"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logger.Logger
	config       *config.Config
	tokenService *TokenService
	db           *gorm.DB
}

// NewUserService creates a new user service
func NewUserService(config *config.Config) *UserService {
	db := db.GetDB()
	logger := logger.NewLogger(config)
	return &UserService{logger: logger, config: config, db: db, tokenService: NewTokenService(config)}
}

// LoginByUsername logs in a user by username
func (s *UserService) LoginByUsername(req *views.LoginByUsernameRequest) (*views.TokenDetail, error) {
	var user models.User
	err := s.db.
		Model(&models.User{}).
		Where("username = ?", req.Username).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	tokenView := tokenDto{UserId: user.Id, Username: user.Username,
		Email: user.Email}

	token, err := s.tokenService.GenerateToken(&tokenView)
	if err != nil {
		return nil, err
	}
	return token, nil

}

// RegisterByUsername registers a user by username
func (s *UserService) RegisterByUsername(req *views.RegisterUserByUsernameRequest) error {
	u := models.User{Username: req.Username, Email: req.Email}

	exists, err := s.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &utils.ServiceError{EndUserMessage: constants.EmailExists}
	}
	exists, err = s.existsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &utils.ServiceError{EndUserMessage: constants.UsernameExists}
	}
	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(logger.General, logger.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)

	tx := s.db.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logger.Postgres, logger.Rollback, err.Error(), nil)
		return err
	}
	if err != nil {
		tx.Rollback()
		s.logger.Error(logger.Postgres, logger.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.db.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logger.Postgres, logger.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.db.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logger.Postgres, logger.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}
