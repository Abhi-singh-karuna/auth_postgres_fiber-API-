package router

import (
	"github.com/abhishek_singh/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/logout", handler.LogOut)


	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

}
