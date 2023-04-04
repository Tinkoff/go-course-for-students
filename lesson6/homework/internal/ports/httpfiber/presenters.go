package httpfiber

import (
	"github.com/gofiber/fiber/v2"

	"homework6/internal/ads"
)

type createAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

type adResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	AuthorID  int64  `json:"author_id"`
	Published bool   `json:"published"`
}

type changeAdStatusRequest struct {
	Published bool  `json:"published"`
	UserID    int64 `json:"user_id"`
}

type updateAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

func AdSuccessResponse(ad *ads.Ad) *fiber.Map {
	return &fiber.Map{
		"data": adResponse{
			ID:        ad.ID,
			Title:     ad.Title,
			Text:      ad.Text,
			AuthorID:  ad.AuthorID,
			Published: ad.Published,
		},
		"error": nil,
	}
}

func AdsSuccessResponse(a []ads.Ad) *fiber.Map {
	var response []adResponse
	for i := range a {
		response = append(response, adResponse{
			ID:        a[i].ID,
			Title:     a[i].Title,
			Text:      a[i].Text,
			AuthorID:  a[i].AuthorID,
			Published: a[i].Published,
		})
	}

	return &fiber.Map{
		"data":  response,
		"error": nil,
	}
}

func AdErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":  nil,
		"error": err.Error(),
	}
}
