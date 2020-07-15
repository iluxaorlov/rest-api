package api

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestApi_HandleHello(t *testing.T) {
	a := New(NewConfig())
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/hello", nil)
	a.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello World")
}