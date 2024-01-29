package services

import (
	"time"

	"github.com/LamichhaneBibek/go-template/internal/config"
	"github.com/LamichhaneBibek/go-template/internal/constants"
	"github.com/LamichhaneBibek/go-template/internal/utils"
	"github.com/LamichhaneBibek/go-template/internal/utils/logger"
	"github.com/LamichhaneBibek/go-template/internal/views"
	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	logger logger.Logger
	config *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(config *config.Config) *TokenService {
	logger := logger.NewLogger(config)
	return &TokenService{
		logger: logger,
		config: config,
	}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*views.TokenDetail, error) {
	td := &views.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(s.config.JWT.AccessExpirationMinutes * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(s.config.JWT.RefreshExpirationMinutes * time.Minute).Unix()

	atc := jwt.MapClaims{}

	atc[constants.UserIdKey] = token.UserId
	atc[constants.FirstNameKey] = token.FirstName
	atc[constants.LastNameKey] = token.LastName
	atc[constants.UsernameKey] = token.Username
	atc[constants.EmailKey] = token.Email
	atc[constants.ExpiredTimeKey] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error
	td.AccessToken, err = at.SignedString([]byte(s.config.JWT.Secret))

	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}

	rtc[constants.UserIdKey] = token.UserId
	rtc[constants.ExpiredTimeKey] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(s.config.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &utils.ServiceError{EndUserMessage: constants.UnExpectedError}
		}
		return []byte(s.config.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &utils.ServiceError{EndUserMessage: constants.ClaimsNotFound}
}
