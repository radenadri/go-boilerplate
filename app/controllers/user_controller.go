package controllers

import (
	pg "boilerplate/app/db"
	"boilerplate/app/models"
	"boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	pg.DB.Find(&users)

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  true,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func GetUser(c *fiber.Ctx) error {
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
