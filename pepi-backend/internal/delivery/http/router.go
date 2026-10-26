// Package http arma la aplicación Fiber: middlewares y rutas.
package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/someonebhguyfvgyfvt6/pepi-backend/internal/delivery/http/handler"
	"github.com/someonebhguyfvgyfvt6/pepi-backend/internal/delivery/http/middleware"
)

// NewRouter crea la app Fiber con todas las rutas registradas.
func NewRouter() *fiber.App {
	app := fiber.New(fiber.Config{AppName: "PEPI API"})

	middleware.Setup(app)

	health := handler.NewHealthHandler()
	app.Get("/health", health.Check)

	return app
}
