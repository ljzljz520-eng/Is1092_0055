package api

import (
	"encoding/json"
	"magazine-editor/model"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func decodeRecord(r *http.Request) (model.Record, error) {
	var v model.Record
	err := json.NewDecoder(r.Body).Decode(&v)
	return v, err
}
func methodAllowed(r *http.Request, allowed ...string) bool {
	for _, m := range allowed {
		if r.Method == m {
			return true
		}
	}
	return false
}
