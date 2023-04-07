package httpfiber

import (
	"net/http"
	"strconv"
	"userservice/internal/app"

	"github.com/gofiber/fiber/v2"
)

func getUser(a app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, _ := strconv.Atoi(c.Params("user_id"))

		u, err := a.GetUser(c.Context(), int64(userID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(UserErrorResponse(err))
		}

		return c.JSON(UserSuccessResponse(u))
	}
}

func createUser(a app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody userNameRequest
		if err := c.BodyParser(&reqBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(UserErrorResponse(err))
		}

		if err := reqBody.validate(); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(UserErrorResponse(err))
		}

		u, err := a.CreateUser(c.Context(), reqBody.Name)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(UserErrorResponse(err))
		}

		return c.JSON(UserSuccessResponse(u))
	}
}
