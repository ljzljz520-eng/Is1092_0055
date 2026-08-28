package editorial

import (
	"context"
	"errors"
	"fmt"
	"magazine-editor/model"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Brief struct {
	Headline, Deck, Angle, Audience string
	Keywords                        []string
}
type Assignment struct {
	ID, RecordID, Editor string
	Brief                Brief
	Due                  time.Time
	State                string
}
type Pipeline struct {
	Assignments map[string]Assignment
	History     map[string][]string
}

func NewPipeline() *Pipeline {
	return &Pipeline{Assignments: map[string]Assignment{}, History: map[string][]string{}}
}
func (p *Pipeline) Assign(a Assignment) error {
	if a.ID == "" || a.RecordID == "" {
		return errors.New("assignment identity required")
	}
	if a.State == "" {
		a.State = "queued"
	}
	p.Assignments[a.ID] = a
	p.History[a.ID] = append(p.History[a.ID], "assigned")
	return nil
}
func (p *Pipeline) Start(id string) error {
	a, ok := p.Assignments[id]
	if !ok {
		return errors.New("assignment missing")
	}
	if a.State != "queued" {
		return fmt.Errorf("cannot start %s", a.State)
	}
	a.State = "active"
	p.Assignments[id] = a
	p.History[id] = append(p.History[id], "started")
	return nil
}
func (p *Pipeline) Complete(id string) error {
	a, ok := p.Assignments[id]
	if !ok {
		return errors.New("assignment missing")
	}
	if a.State != "active" {
		return fmt.Errorf("cannot complete %s", a.State)
	}
	a.State = "complete"
	p.Assignments[id] = a
	p.History[id] = append(p.History[id], "completed")
	return nil
}
func (p *Pipeline) Cancel(id string) error {
	a, ok := p.Assignments[id]
	if !ok {
		return errors.New("assignment missing")
	}
	if a.State == "complete" {
		return errors.New("complete assignment")
	}
	a.State = "cancelled"
	p.Assignments[id] = a
	p.History[id] = append(p.History[id], "cancelled")
	return nil
}
func (p *Pipeline) Due(now time.Time) []Assignment {
	out := []Assignment{}
	for _, a := range p.Assignments {
		if a.State != "complete" && a.State != "cancelled" && !a.Due.After(now) {
			out = append(out, a)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Due.Before(out[j].Due) })
	return out
}
func (p *Pipeline) ForEditor(name string) []Assignment {
	out := []Assignment{}
	for _, a := range p.Assignments {
		if a.Editor == name {
			out = append(out, a)
		}
	}
	return out
}
func (p *Pipeline) Timeline(id string) []string { return append([]string{}, p.History[id]...) }

type Finding struct {
	Code, Message string
	Severity      int
}

func ScanBrief(b Brief) []Finding {
	f := []Finding{}
	if strings.TrimSpace(b.Headline) == "" {
		f = append(f, Finding{"headline.empty", "headline required", 3})
	}
	if len(b.Headline) > 90 {
		f = append(f, Finding{"headline.long", "headline too long", 2})
	}
	if len(strings.TrimSpace(b.Deck)) < 10 {
		f = append(f, Finding{"deck.short", "deck needs context", 1})
	}
	if strings.TrimSpace(b.Angle) == "" {
		f = append(f, Finding{"angle.empty", "angle required", 2})
	}
	if b.Audience == "" {
		f = append(f, Finding{"audience.empty", "audience required", 1})
	}
	if len(b.Keywords) == 0 {
		f = append(f, Finding{"keywords.none", "add keywords", 1})
	}
	return f
}
func ScanBody(body string) []Finding {
	f := []Finding{}
	if strings.TrimSpace(body) == "" {
		return append(f, Finding{"body.empty", "body required", 3})
	}
	if len(body) < 120 {
		f = append(f, Finding{"body.short", "body is too short", 2})
	}
	if !strings.HasSuffix(strings.TrimSpace(body), ".") {
		f = append(f, Finding{"body.ending", "finish with punctuation", 1})
	}
	if strings.Count(body, "\n\n") == 0 {
		f = append(f, Finding{"body.structure", "add paragraph breaks", 1})
	}
	return f
}
func CheckKeywords(text string, words []string) []Finding {
	f := []Finding{}
	low := strings.ToLower(text)
	for _, w := range words {
		if strings.TrimSpace(w) == "" {
			continue
		}
		if !strings.Contains(low, strings.ToLower(w)) {
			f = append(f, Finding{"keyword.missing", w, 1})
		}
	}
	return f
}
func DetectSensitive(text string) []Finding {
	patterns := map[string]string{"phone": `\\b1[3-9]\\d{9}\\b`, `email`: `[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}`, "secret": `(?i)password|token|secret`}
	out := []Finding{}
	for code, pat := range patterns {
		if regexp.MustCompile(pat).MatchString(text) {
			out = append(out, Finding{"sensitive." + code, "sensitive data detected", 3})
		}
	}
	return out
}
func Score(findings []Finding) int {
	score := 100
	for _, f := range findings {
		score -= f.Severity * 10
	}
	if score < 0 {
		return 0
	}
	return score
}
func Ready(findings []Finding) bool {
	for _, f := range findings {
		if f.Severity >= 3 {
			return false
		}
	}
	return true
}
func Review(ctx context.Context, b Brief, body string) (int, []Finding, error) {
	if err := ctx.Err(); err != nil {
		return 0, nil, err
	}
	all := ScanBrief(b)
	all = append(all, ScanBody(body)...)
	all = append(all, CheckKeywords(b.Headline+" "+body, b.Keywords)...)
	all = append(all, DetectSensitive(body)...)
	return Score(all), all, nil
}
func NormalizeBrief(b Brief) Brief {
	b.Headline = strings.TrimSpace(b.Headline)
	b.Deck = strings.TrimSpace(b.Deck)
	b.Angle = strings.TrimSpace(b.Angle)
	b.Audience = strings.TrimSpace(b.Audience)
	clean := []string{}
	seen := map[string]bool{}
	for _, k := range b.Keywords {
		k = strings.ToLower(strings.TrimSpace(k))
		if k != "" && !seen[k] {
			seen[k] = true
			clean = append(clean, k)
		}
	}
	b.Keywords = clean
	return b
}
func WordCount(text string) int { return len(strings.Fields(text)) }
func ReadingMinutes(text string) int {
	n := WordCount(text)
	if n == 0 {
		return 0
	}
	return (n + 199) / 200
}
func DraftSnapshot(r model.Record) map[string]string {
	return map[string]string{"id": r.ID, "title": r.Title, "status": r.Status, "section": r.Section, "version": fmt.Sprint(r.Version)}
}
