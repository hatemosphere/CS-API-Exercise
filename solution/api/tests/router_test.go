package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/hatemosphere/API-Exercise/solution/api/handlers/passengers"
)

func TestList(t *testing.T) {
	router := gin.New()
	router.GET("/testList", passengers.List)

	req, _ := http.NewRequest("GET", "/stuff", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.NotNil(t, resp.Body.String())
}

func TestCreate(t *testing.T) {
	router := gin.New()
	router.GET("/testCreate", passengers.Create)

	json := []byte(`{"animal":"furby"}`)

	req, _ := http.NewRequest("POST", "/stuff", bytes.NewBuffer(json))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.NotNil(t, resp.Body.String())
}
