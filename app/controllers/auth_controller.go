package controllers

import (
	"boilerplate/app/config"
	pg "boilerplate/app/db"
	"boilerplate/app/middlewares"
	"boilerplate/app/models"
	"boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	var user models.User
	if err := pg.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "Invalid email or password",
			Data:    nil,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "Invalid email or password",
			Data:    nil,
		})
	}

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

func Register(c *fiber.Ctx) error {
	registerRequest := new(models.RegisterRequest)

	if err := c.BodyParser(registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not register user",
			Data:    nil,
		})
	}

	user := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: string(hashedPassword),
	}

	if err := pg.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  false,
			Message: "Could not register user",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User registered successfully",
		Data:    user,
	})
}

func Profile(c *fiber.Ctx) error {
	auth := c.Request().Header.Peek("Authorization")

	splitToken := strings.Split(string(auth), "Bearer ")
	auth = []byte(splitToken[1])

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

func UpdateProfile(c *fiber.Ctx) error {
	auth := c.Request().Header.Peek("Authorization")

	splitToken := strings.Split(string(auth), "Bearer ")
	auth = []byte(splitToken[1])

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

	updateUserRequest := new(models.UpdateUserRequest)

	if err := c.BodyParser(updateUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	var user models.User
	if err := pg.DB.Where("id = ?", token.Claims.(jwt.MapClaims)["user_id"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

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
