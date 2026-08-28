package store

import (
	"magazine-editor/model"
	"path/filepath"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "a.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := model.NewRecord("1", "Title", "A sufficiently long body for testing.", "culture")
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetRecord("1"); e != nil {
		t.Fatal(e)
	}
}
