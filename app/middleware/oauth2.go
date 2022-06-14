package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Oauth2Authentication authenticating oauth2 before calling next request
func Oauth2Authentication(c *fiber.Ctx) error {
	// do something here ...

	return c.Next()
}
