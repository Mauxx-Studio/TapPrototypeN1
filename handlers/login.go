package handlers

import (
	"github.com/MauxxStudio/tapirus1go/db"
	"github.com/MauxxStudio/tapirus1go/models"
	"github.com/MauxxStudio/tapirus1go/session"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var user models.User
	var userlogin models.UserRegister

	if err := c.BodyParser(&userlogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	result := db.DB.Where("email = ?", userlogin.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userlogin.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := session.GenerateToken(user.ID.String())
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Usuario no autorizado"})
	}

	c.Set("Authorization", token)

	return c.JSON("Usuario logueado")
}
