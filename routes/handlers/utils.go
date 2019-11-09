package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseReqJsonBody(body interface{}, w http.ResponseWriter, r *http.Request) {
	buffer, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(buffer, body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Status(status HttpStatus) int {
	return int(status)
}

func String(s string) *string {
	return &s
}
