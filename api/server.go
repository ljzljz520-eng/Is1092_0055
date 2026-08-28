package api

import (
	"encoding/json"
	"magazine-editor/model"
	"magazine-editor/service"
	"net/http"
)

type Server struct{ Editor *service.Editor }

func (s Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200); w.Write([]byte("ok")) })
	m.HandleFunc("/records", s.records)
	return m
}
func (s Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var rec model.Record
		if json.NewDecoder(r.Body).Decode(&rec) != nil {
			http.Error(w, "bad json", 400)
			return
		}
		if err := s.Editor.Register(r.Context(), rec); err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		json.NewEncoder(w).Encode(rec)
		return
	}
	term := r.URL.Query().Get("q")
	status := r.URL.Query().Get("status")
	rs, err := s.Editor.Query(term, status)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(rs)
}
