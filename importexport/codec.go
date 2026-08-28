package importexport

import (
	"encoding/json"
	"magazine-editor/model"
)

func EncodeRecord(r model.Record) ([]byte, error) { return json.Marshal(r) }
func DecodeRecord(b []byte) (model.Record, error) {
	var r model.Record
	err := json.Unmarshal(b, &r)
	return r, err
}
func EncodeMany(rs []model.Record) ([]byte, error) { return json.Marshal(rs) }
func DecodeMany(b []byte) ([]model.Record, error) {
	var rs []model.Record
	err := json.Unmarshal(b, &rs)
	return rs, err
}
