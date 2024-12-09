package handlers

import (
	"github.com/MauxxStudio/tapirus1go/db"
	"github.com/MauxxStudio/tapirus1go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	query := c.Query("name")
	query = "%" + query + "%"

	if query == "" {
		db.DB.Limit(5).Find(&users)
	} else {
		result := db.DB.Limit(5).Where("first_name LIKE ? OR last_name LIKE ?", query, query).Find(&users)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
	}

	var userts models.UserToSend
	var usersts []models.UserToSend

	for _, user := range users {
		userts.FirstName = user.FirstName
		userts.LastName = user.LastName
		userts.Email = user.Email
		userts.ID = user.ID
		usersts = append(usersts, userts)
	}

	return c.Status(fiber.StatusOK).JSON(usersts)
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	param := c.Params("id")

	userFinded := db.DB.First(&user, param)

	if userFinded.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func NewUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}
	if len(user.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "password is too weak"})
	}

	user.ID = uuid.New()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Hash Fail"})
	}
	user.Password = string(hash)

	createdUser := db.DB.Create(&user)
	err = createdUser.Error
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": "invalid user"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User

	param := c.Params("id")

	userFinded := db.DB.First(&user, param)

	if userFinded.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "invalid id"})
	}

	userDeleted := db.DB.Delete(&user)
	if userDeleted.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "delete fail"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
