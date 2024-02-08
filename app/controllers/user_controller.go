package controllers

import (
	pg "boilerplate/app/db"
	"boilerplate/app/models"
	"boilerplate/app/utils"

	"github.com/gofiber/fiber/v2"
)

// GetUsers godoc
// @Summary      List all users
// @Description  Get a list of all users
// @Tags         User
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /users [get]
func GetUsers(c *fiber.Ctx) error {
	// Get users
	var users []models.User
	pg.DB.Find(&users)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

// GetUser godoc
// @Summary      Get user
// @Description  Get user by ID
// @Tags         User
// @Param        id   path      int  true  "User ID"
// @Accept       json
// @Headers      Content-Type application/json
// @Produce      json
// @Security	 ApiKeyAuth
// @Success      200  {array}   models.User
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	// Get user
	var user models.User
	if err := pg.DB.First(&user, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}
