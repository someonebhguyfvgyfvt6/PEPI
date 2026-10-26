// Package middleware registra los middlewares globales de la API.
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Setup aplica los middlewares comunes a todas las rutas.
func Setup(app *fiber.App) {
	app.Use(recover.New()) // evita que un panic tumbe el servidor
	app.Use(logger.New())  // registra cada petición en consola
}
