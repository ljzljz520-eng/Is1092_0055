package catalog

import (
	"magazine-editor/model"
	"testing"
)

func TestCatalogSearch(t *testing.T) {
	c := New()
	_ = c.Add(model.NewRecord("1", "Culture", "A long body for catalog test.", "culture"))
	if len(c.Search("cult")) != 1 {
		t.Fatal()
	}
}
