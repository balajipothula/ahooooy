package http

import (
    "dating-app/services/login/service"
    "github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
    Auth *service.AuthService
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
    }

    user, err := h.Auth.Register(req.Email, req.Password)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
    }

    token, err := h.Auth.Login(req.Email, req.Password)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
    }

    return c.JSON(fiber.Map{"token": token})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
    token := c.Get("Authorization")
    if token == "" {
        return c.Status(401).JSON(fiber.Map{"error": "missing token"})
    }

    if err := h.Auth.Logout(token); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "logged out"})
}
