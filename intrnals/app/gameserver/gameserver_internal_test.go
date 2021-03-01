package gameserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameServer_handleTest(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)

	s.handleTest().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "Test Request!")
}
