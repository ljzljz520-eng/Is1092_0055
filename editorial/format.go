package editorial

import (
	"regexp"
	"strings"
)

var spaces = regexp.MustCompile(`\\s+`)

func CleanText(v string) string { return strings.TrimSpace(spaces.ReplaceAllString(v, " ")) }
func Sentences(v string) []string {
	parts := strings.FieldsFunc(v, func(r rune) bool { return r == '.' || r == '!' || r == '?' })
	o := []string{}
	for _, p := range parts {
		if s := CleanText(p); s != "" {
			o = append(o, s)
		}
	}
	return o
}
func Truncate(v string, n int) string {
	r := []rune(v)
	if len(r) <= n {
		return v
	}
	return string(r[:n]) + "..."
}
func Slug(v string) string { return strings.ToLower(strings.ReplaceAll(CleanText(v), " ", "-")) }
