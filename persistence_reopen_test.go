package magazine_editor

import (
	"magazine-editor/model"
	"magazine-editor/store"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "persist.db")
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.NewRecord("persist", "Title", "A sufficiently long body for testing.", "culture")
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord("persist"); e != nil {
		t.Fatal(e)
	}
}
