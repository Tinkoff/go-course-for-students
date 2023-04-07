package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"homework6/internal/adapters/adrepo"
	"homework6/internal/app"
	"homework6/internal/ports/httpfiber"
)

type adData struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	AuthorID  int64  `json:"author_id"`
	Published bool   `json:"published"`
}

type adResponse struct {
	Data adData `json:"data"`
}

var ErrBadRequest = fmt.Errorf("bad request")
var ErrForbidden = fmt.Errorf("forbidden")

func getResponse(server httpfiber.Server, req *http.Request, out interface{}) error {
	resp, err := server.Test(req)
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return ErrBadRequest
		}
		if resp.StatusCode == http.StatusForbidden {
			return ErrForbidden
		}
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}

func createAd(server httpfiber.Server, userID int64, title string, text string) (adResponse, error) {
	body := map[string]any{
		"user_id": userID,
		"title":   title,
		"text":    text,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return adResponse{}, fmt.Errorf("unable to marshal: %w", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/ads", bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	var response adResponse
	err = getResponse(server, req, &response)
	if err != nil {
		return adResponse{}, err
	}

	return response, nil
}

func changeAdStatus(server httpfiber.Server, userID int64, adID int64, published bool) (adResponse, error) {
	body := map[string]any{
		"user_id":   userID,
		"published": published,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return adResponse{}, fmt.Errorf("unable to marshal: %w", err)
	}

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/ads/%d/status", adID), bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	var response adResponse
	err = getResponse(server, req, &response)
	if err != nil {
		return adResponse{}, err
	}

	return response, nil
}

func updateAd(server httpfiber.Server, userID int64, adID int64, title string, text string) (adResponse, error) {
	body := map[string]any{
		"user_id": userID,
		"title":   title,
		"text":    text,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return adResponse{}, fmt.Errorf("unable to marshal: %w", err)
	}

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/ads/%d", adID), bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	var response adResponse
	err = getResponse(server, req, &response)
	if err != nil {
		return adResponse{}, err
	}

	return response, nil
}

func TestCreateAd(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	response, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if response.Data.ID != 0 {
		t.Errorf("invalid id")
	}
	if response.Data.Title != "hello" {
		t.Errorf("invalid title")
	}
	if response.Data.Text != "world" {
		t.Errorf("invalid text")
	}
	if response.Data.AuthorID != 123 {
		t.Errorf("invalid author id")
	}
	if response.Data.Published {
		t.Errorf("invalid published state")
	}
}

func TestChangeAdStatus(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	response, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	response, err = changeAdStatus(server, 123, response.Data.ID, true)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !response.Data.Published {
		t.Errorf("invalid published status")
	}

	response, err = changeAdStatus(server, 123, response.Data.ID, false)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if response.Data.Published {
		t.Errorf("invalid published status")
	}

	// Check that unpublishing unpublished does nothing

	response, err = changeAdStatus(server, 123, response.Data.ID, false)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if response.Data.Published {
		t.Errorf("invalid published status")
	}
}

func TestUpdateAd(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	response, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	response, err = updateAd(server, 123, response.Data.ID, "привет", "мир")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if response.Data.Title != "привет" {
		t.Errorf("invalid title")
	}

	if response.Data.Text != "мир" {
		t.Errorf("invalid text")
	}
}
