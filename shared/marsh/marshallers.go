package marsh

import (
	"encoding/json"
	http2 "github.com/evscott/z3-e2c-api/shared/http"
	"io/ioutil"
	"log"
	"net/http"
)

func UnmarshalRequest(body interface{}, w http.ResponseWriter, r *http.Request) {
	buffer, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http2.Status(http2.InternalServerError))
		log.Fatal(err)
	}

	err = json.Unmarshal(buffer, body)
	if err != nil {
		http.Error(w, err.Error(), http2.Status(http2.InternalServerError))
		log.Fatal(err)
	}
}

func MarshalResponse(body interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		http.Error(w, err.Error(), http2.Status(http2.InternalServerError))
		log.Fatal(err)
	}
}
