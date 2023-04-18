package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAd_EmptyTitle(t *testing.T) {
	client := getTestClient()

	_, err := client.createAd(123, "", "world")
	assert.ErrorIs(t, err, ErrBadRequest)
}

func TestCreateAd_TooLongTitle(t *testing.T) {
	client := getTestClient()

	title := strings.Repeat("a", 101)

	_, err := client.createAd(123, title, "world")
	assert.ErrorIs(t, err, ErrBadRequest)
}

func TestCreateAd_EmptyText(t *testing.T) {
	client := getTestClient()

	_, err := client.createAd(123, "title", "")
	assert.ErrorIs(t, err, ErrBadRequest)
}

func TestCreateAd_TooLongText(t *testing.T) {
	client := getTestClient()

	text := strings.Repeat("a", 501)

	_, err := client.createAd(123, "title", text)
	assert.ErrorIs(t, err, ErrBadRequest)
}

func TestUpdateAd_EmptyTitle(t *testing.T) {
	client := getTestClient()

	resp, err := client.createAd(123, "hello", "world")
	assert.NoError(t, err)

	_, err = client.updateAd(123, resp.Data.ID, "", "new_world")
	assert.ErrorIs(t, err, ErrBadRequest)
}

func TestUpdateAd_TooLongTitle(t *testing.T) {
	client := getTestClient()

	resp, err := client.createAd(123, "hello", "world")
	assert.NoError(t, err)

	title := strings.Repeat("a", 101)

	_, err = client.updateAd(123, resp.Data.ID, title, "world")
	assert.ErrorIs(t, err, ErrBadRequest)
}

func TestUpdateAd_EmptyText(t *testing.T) {
	client := getTestClient()

	resp, err := client.createAd(123, "hello", "world")
	assert.NoError(t, err)

	_, err = client.updateAd(123, resp.Data.ID, "title", "")
	assert.ErrorIs(t, err, ErrBadRequest)
}

func TestUpdateAd_TooLongText(t *testing.T) {
	client := getTestClient()

	text := strings.Repeat("a", 501)

	resp, err := client.createAd(123, "hello", "world")
	assert.NoError(t, err)

	_, err = client.updateAd(123, resp.Data.ID, "title", text)
	assert.ErrorIs(t, err, ErrBadRequest)
}
