package model

import "strings"

func NormalizeTitle(v string) string { return strings.TrimSpace(strings.Join(strings.Fields(v), " ")) }
func AllowedStatus(v string) bool {
	switch v {
	case "draft", "review", "approved", "published", "archived", "rejected":
		return true
	}
	return false
}
func ValidSection(v string) bool {
	if v == "" {
		return false
	}
	return len(v) <= 64
}
func (r Record) Summary() string {
	if len(r.Body) > 120 {
		return r.Body[:120]
	}
	return r.Body
}
func (p Profile) Label() string { return p.Name + " (" + p.Role + ")" }
