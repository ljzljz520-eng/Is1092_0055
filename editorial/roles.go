package editorial

import (
	"sort"
	"strings"
)

type Role struct {
	Name        string
	Permissions []string
}

func Roles() []Role {
	return []Role{{"editor", []string{"edit", "submit", "view"}}, {"chief", []string{"edit", "submit", "review", "publish", "archive"}}, {"reviewer", []string{"view", "review"}}}
}
func HasPermission(role, perm string) bool {
	for _, r := range Roles() {
		if r.Name == role {
			for _, p := range r.Permissions {
				if p == perm {
					return true
				}
			}
		}
	}
	return false
}
func RoleNames() []string {
	o := []string{}
	for _, r := range Roles() {
		o = append(o, r.Name)
	}
	sort.Strings(o)
	return o
}
func NormalizeRole(r string) string { return strings.ToLower(strings.TrimSpace(r)) }
