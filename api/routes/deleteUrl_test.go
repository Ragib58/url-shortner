package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteURL_NotFound(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.DELETE("/:shortID", DeleteURL)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/nonexistent", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Add more tests for other routes as needed
