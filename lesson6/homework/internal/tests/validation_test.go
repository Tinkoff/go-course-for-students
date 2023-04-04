package tests

import (
	"errors"
	"strings"
	"testing"

	"homework6/internal/adapters/adrepo"
	"homework6/internal/app"
	"homework6/internal/ports/httpfiber"
)

func TestCreateAd_EmptyTitle(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	_, err := createAd(server, 123, "", "world")
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("expected error")
	}
}

func TestCreateAd_TooLongTitle(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	title := strings.Repeat("a", 101)

	_, err := createAd(server, 123, title, "world")
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("expected error")
	}
}

func TestCreateAd_EmptyText(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	_, err := createAd(server, 123, "title", "")
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("expected error")
	}
}

func TestCreateAd_TooLongText(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	text := strings.Repeat("a", 501)

	_, err := createAd(server, 123, "title", text)
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("expected error")
	}
}

func TestUpdateAd_EmptyTitle(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	resp, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	_, err = updateAd(server, 123, resp.Data.ID, "", "new_world")
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("expected error")
	}
}

func TestUpdateAd_TooLongTitle(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	resp, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	title := strings.Repeat("a", 101)

	_, err = updateAd(server, 123, resp.Data.ID, title, "world")
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("expected error")
	}
}

func TestUpdateAd_EmptyText(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	resp, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	_, err = updateAd(server, 123, resp.Data.ID, "title", "")
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("expected error")
	}
}

func TestUpdateAd_TooLongText(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	text := strings.Repeat("a", 501)

	resp, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	_, err = updateAd(server, 123, resp.Data.ID, "title", text)
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("expected error")
	}
}
