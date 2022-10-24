package service

import "github.com/gofiber/fiber/v2"

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Users APIs
	api.Post("/create_users", r.CreateUser)
	api.Delete("/delete_user/:vr_app_session", r.DeleteUser)
	api.Put("/update_user/:vr_app_session", r.UpdateUser)
	api.Get("/get_users/:vr_app_session", r.GetUserByID)
	api.Get("/users", r.GetUsers)

	// News APIs
	api.Post("/create_news", r.CreateNews)
	api.Delete("/delete_news/:news_id", r.DeleteNews)
	api.Put("/update_news/:news_id", r.UpdateNews)
	api.Get("/news", r.GetNews)
}