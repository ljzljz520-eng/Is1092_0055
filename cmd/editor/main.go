package main

import (
	"log"
	"magazine-editor/api"
	"magazine-editor/service"
	"magazine-editor/store"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("EDITOR_DB")
	if path == "" {
		path = "editor.db"
	}
	s, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	log.Println(http.ListenAndServe(":8080", api.Server{Editor: &service.Editor{Store: s}}.Handler()))
}
