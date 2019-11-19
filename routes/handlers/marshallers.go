package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func UnmarshalRequest(body interface{}, w http.ResponseWriter, r *http.Request) {
	buffer, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), Status(InternalServerError))
		log.Fatal(err)
	}

	err = json.Unmarshal(buffer, body)
	if err != nil {
		http.Error(w, err.Error(), Status(InternalServerError))
		log.Fatal(err)
	}
}

func MarshalResponse(body interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		http.Error(w, err.Error(), Status(InternalServerError))
		log.Fatal(err)
	}
}
