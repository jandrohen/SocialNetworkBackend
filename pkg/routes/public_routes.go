package routes

import (
	"SocialNetworkBackend/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// // Routes for POST method:
	route.Post("/user/sign/in", controllers.UserSignIn)  // auth, return Access & Refresh tokens

	route.Post("/user/sign/up", controllers.UserSignUp) // register a new user

	// Create routes group.
	route2 := a.Group("/")

	// Routes for GET method:
	route2.Static("/app", "./public")

}