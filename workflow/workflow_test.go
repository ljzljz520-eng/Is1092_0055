package workflow

import (
	"context"
	"magazine-editor/model"
	"magazine-editor/service"
	"magazine-editor/store"
	"path/filepath"
	"testing"
)

func setup(t *testing.T) Chain {
	s, e := store.Open(filepath.Join(t.TempDir(), "w.db"))
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return Chain{Editor: &service.Editor{Store: s}}
}
func TestWorkflowOne(t *testing.T) {
	c := setup(t)
	r := model.NewRecord("one", "Feature", "This is a long enough feature body for tests.", "culture")
	if e := c.Receive(context.Background(), r); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	c := setup(t)
	r := model.NewRecord("two", "Feature", "This is a long enough feature body for tests.", "culture")
	if e := c.Full(context.Background(), r); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	c := setup(t)
	r := model.NewRecord("three", "Feature", "This is a long enough feature body for tests.", "culture")
	if e := c.Receive(context.Background(), r); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain27(t *testing.T) {
	c := setup(t)
	r := model.NewRecord("bug", "Feature", "short", "culture")
	r.Status = "approved"
	if e := c.Editor.Store.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if e := c.Editor.Publish(context.Background(), r.ID); e == nil {
		t.Fatal("invalid draft must not publish")
	}
}
