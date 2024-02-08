package controllers

import (
	"boilerplate/app/config"
	pg "boilerplate/app/db"
	"boilerplate/app/middlewares"
	"boilerplate/app/models"
	"boilerplate/app/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary      Perform login
// @Description  Login with email and password
// @Tags         Auth
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Param        request body models.LoginRequest true "Login request"
// @Success      200  {array}   models.LoginResponse
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /login [post]
func Login(c *fiber.Ctx) error {
	// Get and parse user input
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	// Validate user input
	validationErrors := utils.GlobalValidator.Validate(loginRequest)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Validation errors",
			Data:    validationErrors,
		})
	}

	// Check if user exists
	var user models.User
	if err := pg.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "Invalid email or password",
			Data:    nil,
		})
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "Invalid email or password",
			Data:    nil,
		})
	}

	// Generate JWT
	token, err := middlewares.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not login user",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User logged in successfully",
		Data:    token,
	})
}

// Register godoc
// @Summary      Attempt register
// @Description  Register to the system
// @Tags         Auth
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Param        request body models.RegisterRequest true "Register request"
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /register [post]
func Register(c *fiber.Ctx) error {
	// Get and parse user input
	registerRequest := new(models.RegisterRequest)
	if err := c.BodyParser(registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	// Validate user input
	validationErrors := utils.GlobalValidator.Validate(registerRequest)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Validation errors",
			Data:    validationErrors,
		})
	}

	// Generate password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not register user",
			Data:    nil,
		})
	}

	// Create user
	user := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: string(hashedPassword),
	}

	// Save user
	if err := pg.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not save user",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User registered successfully",
		Data:    user,
	})
}

// Profile godoc
// @Summary      Get profile
// @Description  Get current user profile
// @Tags         Auth
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /me [get]
func Profile(c *fiber.Ctx) error {
	// Get header "Authorization"
	auth := c.Request().Header.Peek("Authorization")

	// Split token
	splitToken := strings.Split(string(auth), "Bearer ")
	auth = []byte(splitToken[1])

	// Parse token
	token, err := jwt.ParseWithClaims(string(auth), jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Error while parsing token",
			Data:    err.Error(),
		})
	}

	// Get user
	var user models.User
	if err := pg.DB.Where("id = ?", token.Claims.(jwt.MapClaims)["user_id"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User profile retrieved successfully",
		Data:    user,
	})
}

// UpdateProfile godoc
// @Summary      Update profile
// @Description  Update current user profile
// @Tags         Auth
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Security	 ApiKeyAuth
// @Param        request body models.UpdateUserRequest true "Update profile request"
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /me [put]
func UpdateProfile(c *fiber.Ctx) error {
	// Get header "Authorization"
	auth := c.Request().Header.Peek("Authorization")

	// Split token
	splitToken := strings.Split(string(auth), "Bearer ")
	auth = []byte(splitToken[1])

	// Parse token
	token, err := jwt.ParseWithClaims(string(auth), jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Error while parsing token",
			Data:    err.Error(),
		})
	}

	// Get user input
	updateUserRequest := new(models.UpdateUserRequest)
	if err := c.BodyParser(updateUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	// Validate user input
	validationErrors := utils.GlobalValidator.Validate(updateUserRequest)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Validation errors",
			Data:    validationErrors,
		})
	}

	// Get user
	var user models.User
	if err := pg.DB.Where("id = ?", token.Claims.(jwt.MapClaims)["user_id"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	// Update user
	user.Name = updateUserRequest.Name
	if err := pg.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Error while updating user",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User profile updated successfully",
		Data:    user,
	})
}
