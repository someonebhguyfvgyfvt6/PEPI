// Package handler contiene los handlers HTTP: traducen peticiones a llamadas de usecases.
package handler

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Check responde 200 si el servidor está vivo.
func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
