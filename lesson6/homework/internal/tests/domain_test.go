package tests

import (
	"errors"
	"testing"

	"homework6/internal/adapters/adrepo"
	"homework6/internal/app"
	"homework6/internal/ports/httpfiber"
)

func TestChangeStatusAdOfAnotherUser(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	resp, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	_, err = changeAdStatus(server, 100, resp.Data.ID, true)
	if !errors.Is(err, ErrForbidden) {
		t.Fatalf("expected error, got: %s", err)
	}
}

func TestUpdateAdOfAnotherUser(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	resp, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	_, err = updateAd(server, 100, resp.Data.ID, "title", "text")
	if !errors.Is(err, ErrForbidden) {
		t.Fatalf("expected error, got: %s", err)
	}
}

func TestCreateAd_ID(t *testing.T) {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))

	respOne, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if respOne.Data.ID != 0 {
		t.Errorf("invalid id")
	}

	respTwo, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if respTwo.Data.ID != 1 {
		t.Errorf("invalid id")
	}

	respThree, err := createAd(server, 123, "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if respThree.Data.ID != 2 {
		t.Errorf("invalid id")
	}

}
