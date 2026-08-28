package audit

import (
	"fmt"
	"magazine-editor/model"
	"magazine-editor/store"
	"time"
)

type Logger struct{ Store *store.Store }

func (l Logger) Record(action, target, actor, detail string) error {
	a := model.Audit{ID: fmt.Sprintf("%d-%s", time.Now().UnixNano(), target), Action: action, Target: target, Actor: actor, Detail: detail, At: time.Now().UTC()}
	return l.Store.SaveAudit(a)
}
func (l Logger) Read(id string) (model.Audit, error) { return l.Store.GetAudit(id) }
func Event(id, rid, kind, actor, payload string) model.Event {
	return model.Event{ID: id, RecordID: rid, Kind: kind, Actor: actor, Payload: payload, At: time.Now().UTC()}
}
