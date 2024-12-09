package handlers

import (
	"github.com/MauxxStudio/tapirus1go/db"
	"github.com/MauxxStudio/tapirus1go/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) error {
	var people []models.Person

	if err := db.DB.Find(&people); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Búsqueda fallida"})
	}

	return c.Status(fiber.StatusOK).JSON(people)
}

func NewPerson(c *fiber.Ctx) error {
	var person models.Person

	if err := c.BodyParser(&person); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	validated := validator.New()
	if err := validated.Struct(person); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	if err := db.DB.Create(&person); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": "Datos incorrectos"})
	}
	return c.Status(fiber.StatusOK).JSON(person)
}

func DeletePerson(c *fiber.Ctx) error {
	var person models.Person
	param := c.Params("id")
	if err := db.DB.First(&person, param); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Person not found"})
	}

	if err := db.DB.Delete(&person); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Person delete failed"})
	}
	return c.Status(fiber.StatusOK).JSON(person)
}
