package httpfiber

import (
	"github.com/gofiber/fiber/v2"

	"homework6/internal/app"
)

func AppRouter(r fiber.Router, a app.App) {
	r.Post("/ads", createAd(a))                    // Метод для создания объявления (ad)
	r.Put("/ads/:ad_id/status", changeAdStatus(a)) // Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
	r.Put("/ads/:ad_id", updateAd(a))              // Метод для обновления текста(Text) или заголовка(Title) объявления
}
