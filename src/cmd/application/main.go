package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vashp/ms-user/src/cmd/infrastructure/connection"
	"github.com/vashp/ms-user/src/cmd/infrastructure/services"

	"github.com/vashp/ms-user/src/cmd/infrastructure/repositories"
	"github.com/vashp/ms-user/src/cmd/presentation/controllers"
	"log"
)

func main() {
	log.Fatal(run())
}

func run() error {
	db, err := connection.ConnectPostgres()
	if err != nil {
		return err
	}

	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)
	controller := controllers.NewController(userService)

	app := fiber.New()
	app.Post("/users", controller.CreateUser)
	app.Get("/users/:id", controller.GetUser)
	app.Put("/users/:id", controller.UpdateUser)
	app.Delete("/users/:id", controller.DeleteUser)

	return app.Listen(":8181")
}
