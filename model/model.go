package model

import "time"

type Record struct {
	ID, Title, Body, Status, Section string
	Version                          int
	CreatedAt, UpdatedAt             time.Time
}
type Profile struct {
	ID, Name, Role string
	Active         bool
	Preferences    map[string]string
}
type Event struct {
	ID, RecordID, Kind, Actor string
	At                        time.Time
	Payload                   string
}
type Audit struct {
	ID, Action, Target, Actor, Detail string
	At                                time.Time
}

func NewRecord(id, title, body, section string) Record {
	now := time.Now().UTC()
	return Record{ID: id, Title: title, Body: body, Section: section, Status: "draft", Version: 1, CreatedAt: now, UpdatedAt: now}
}
func (r Record) IsEditable() bool    { return r.Status == "draft" || r.Status == "rejected" }
func (r Record) IsPublishable() bool { return r.Status == "approved" }
func (r Record) WithStatus(s string) Record {
	r.Status = s
	r.Version++
	r.UpdatedAt = time.Now().UTC()
	return r
}
func (p Profile) Can(action string) bool {
	if !p.Active {
		return false
	}
	if p.Role == "editor" || p.Role == "chief" {
		return true
	}
	return action == "view"
}
func (e Event) Valid() bool { return e.ID != "" && e.RecordID != "" && e.Kind != "" }
func (a Audit) Valid() bool { return a.ID != "" && a.Action != "" && a.Target != "" }
