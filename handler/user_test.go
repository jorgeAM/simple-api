package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jorgeAM/api/models"
)

func TestGet(t *testing.T) {
	h := &Handler{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/test", nil)
	h.Get(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Codigo de estado %d, se obtuvo: %d", http.StatusOK, w.Code)
	}

	u := new(models.User)
	json.NewDecoder(w.Body).Decode(u)

	t.Log(u)
}
