package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jorgeAM/api/models"
)

//DisplayMessage show message with http code
func DisplayMessage(w http.ResponseWriter, m *models.Response) {
	bytes, err := json.Marshal(m)

	if err != nil {
		log.Fatal("something got wrong to convert to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m.Code)
	w.Write(bytes)
}
