package api

import (
	"magazine-editor/service"
	"magazine-editor/store"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "a.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	Server{Editor: &service.Editor{Store: s}}.Handler().ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
func TestCreateRecord(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "a.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	body := `{"id":"a","title":"Feature","body":"This is a long enough feature body for tests.","section":"culture","status":"draft"}`
	req := httptest.NewRequest("POST", "/records", strings.NewReader(body))
	w := httptest.NewRecorder()
	Server{Editor: &service.Editor{Store: s}}.Handler().ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
