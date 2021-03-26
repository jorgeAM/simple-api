package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func DisplayMessage(w http.ResponseWriter, r *Response) {
	bytes, err := json.Marshal(r)

	if err != nil {
		log.Fatal("something got wrong to convert to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	w.Write(bytes)
}
