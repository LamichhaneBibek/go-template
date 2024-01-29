package handlers

import (
	"net/http"

	"github.com/LamichhaneBibek/go-template/internal/api/helpers"
	"github.com/LamichhaneBibek/go-template/internal/config"
	"github.com/LamichhaneBibek/go-template/internal/constants"
	"github.com/LamichhaneBibek/go-template/internal/services"
	"github.com/LamichhaneBibek/go-template/internal/views"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.UserService
}

func NewUsersHandler(config *config.Config) *AuthHandler {
	service := services.NewUserService(config)
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	req := new(views.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, constants.ValidationError, err))
		return
	}
	token, err := h.service.LoginByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, constants.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(token, true, constants.Success))
}

func (h *AuthHandler) Register(c *gin.Context) {
	req := new(views.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, constants.ValidationError, err))
		return
	}
	err = h.service.RegisterByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, constants.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(nil, true, constants.Success))
}
