package importexport

import (
	"magazine-editor/model"
	"testing"
)

func TestCodec(t *testing.T) {
	r := model.NewRecord("1", "Title", "A sufficiently long body for testing.", "culture")
	b, e := EncodeRecord(r)
	if e != nil {
		t.Fatal(e)
	}
	if _, e = DecodeRecord(b); e != nil {
		t.Fatal(e)
	}
}
