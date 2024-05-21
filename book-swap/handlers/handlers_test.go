package handlers_test

import (
	"book-swap/db"
	"book-swap/handlers"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestIndexIntegration in short mode.")
	}

	// Arrange
	book := db.Book{
		ID:     uuid.New().String(),
		Name:   "My first intergration test",
		Status: db.Available.String(),
	}
	bs := db.NewBookService([]db.Book{book}, nil)
	h := handlers.NewHandler(bs, nil)
	svr := httptest.NewServer(http.HandlerFunc(h.Index))
	defer svr.Close()

	// Act
	r, err := http.Get(svr.URL)

	// Assert
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	require.Nil(t, err)
	var resp handlers.Response
	err = json.Unmarshal(body, &resp)
	require.Nil(t, err)
	assert.Equal(t, 1, len(resp.Books))
	assert.Contains(t, resp.Books, book)
}
