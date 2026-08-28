package service

import (
	"context"
	"magazine-editor/model"
	"magazine-editor/store"
	"path/filepath"
	"testing"
)

func testEditor(t *testing.T) *Editor {
	s, e := store.Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return &Editor{Store: s}
}
func TestEditorLifecycle(t *testing.T) {
	e := testEditor(t)
	r := model.NewRecord("r", "Feature", "This is a long enough feature body for tests.", "culture")
	if err := e.Register(context.Background(), r); err != nil {
		t.Fatal(err)
	}
	if err := e.Submit(context.Background(), "r"); err != nil {
		t.Fatal(err)
	}
	if err := e.Review(context.Background(), "r", true); err != nil {
		t.Fatal(err)
	}
}
