package marsh

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UnmarshalRequest(body interface{}, w http.ResponseWriter, r *http.Request) error {
	buffer, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(buffer, body)
}

func MarshalResponse(body interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(body)
}
