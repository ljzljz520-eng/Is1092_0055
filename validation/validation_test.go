package validation

import (
	"magazine-editor/model"
	"testing"
)

func TestValidateRecord(t *testing.T) {
	if ValidateRecord(model.NewRecord("1", "x", "short", "culture")).Valid() {
		t.Fatal("expected issues")
	}
}
