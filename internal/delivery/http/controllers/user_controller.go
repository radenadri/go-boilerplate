package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/radenadri/go-boilerplate/config"
	"github.com/radenadri/go-boilerplate/internal/delivery/dto/request"
	"github.com/radenadri/go-boilerplate/internal/delivery/dto/response"
	"github.com/radenadri/go-boilerplate/internal/domain/models"
	"github.com/radenadri/go-boilerplate/internal/repositories"
	"github.com/radenadri/go-boilerplate/internal/services"
	"github.com/radenadri/go-boilerplate/pkg"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController() UserController {
	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)

	return UserController{
		UserService: userService,
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users with pagination
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/users [get]
func (controller *UserController) GetAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	results, err := controller.UserService.GetAllUsers(page, perPage)

	// Logging example using zap
	config.Logger.Info("Get all users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User registration information"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/v1/register [post]
func (controller *UserController) Register(c *gin.Context) {
	var userPayload models.User

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	if err := pkg.Validator.Struct(userPayload); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Errors:  pkg.FormatValidationErrors(err),
		})
		return
	}

	userResponse, err := controller.UserService.Register(userPayload)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Success: true,
		Data:    userResponse,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body request.UserLoginRequest true "Login credentials"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/v1/login [post]
func (controller *UserController) Login(c *gin.Context) {
	var userLoginPayload request.UserLoginRequest

	if err := c.ShouldBindJSON(&userLoginPayload); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	if err := pkg.Validator.Struct(userLoginPayload); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Errors:  pkg.FormatValidationErrors(err),
		})
		return
	}

	userResponse, err := controller.UserService.Login(userLoginPayload)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Success: true,
		Data:    userResponse,
	})
}
