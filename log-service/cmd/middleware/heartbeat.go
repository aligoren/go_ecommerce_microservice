package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func HeartBeat(endpoint string) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		methodName := ctx.Method()
		path := ctx.Path()

		if (methodName == "GET" || methodName == "HEAD") && strings.EqualFold(path, endpoint) {
			return ctx.Status(http.StatusOK).SendString(".")
		}

		return ctx.Next()
	}
}
