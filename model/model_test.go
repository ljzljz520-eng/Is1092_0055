package model

import "testing"

func TestRecordPredicates(t *testing.T) {
	r := NewRecord("1", "Title", "A sufficiently long body for testing.", "culture")
	if !r.IsEditable() {
		t.Fatal()
	}
	if r.WithStatus("approved").IsPublishable() == false {
		t.Fatal()
	}
}
