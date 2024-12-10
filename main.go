package main

import (
	"fmt"
	"log"

	"github.com/MauxxStudio/tapirus1go/db"
	"github.com/MauxxStudio/tapirus1go/handlers"
	"github.com/MauxxStudio/tapirus1go/models"
	"github.com/gofiber/fiber/v2"
	//	"github.com/gofiber/fiber/v2/middleware/logger"
	//	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	db.DBConnection()

	if err := db.DB.AutoMigrate(models.User{}); err != nil {
		fmt.Println("table users not created")
	}
	if err := db.DB.AutoMigrate(models.Person{}); err != nil {
		fmt.Println("table people not created")
	}

	app := fiber.New()
	//	app.Use(logger.New())
	//	store := session.New()
	//	app.Use(store)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// Users routes
	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.NewUser)
	app.Delete("users/:id", handlers.DeleteUser)

	app.Post("/login", handlers.Login)

	//Staff routes
	app.Get("/staff", handlers.GetPerson)
	app.Post("/staff", handlers.NewPerson)
	app.Delete("/staff/:id", handlers.DeletePerson)

	//Attendance routes

	log.Fatal(app.Listen(":3000"))
}
