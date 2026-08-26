package editorial

import (
	"fmt"
	"time"
)

type Notice struct {
	ID, Recipient, Subject, Body string
	SentAt                       time.Time
	Read                         bool
}

func NewNotice(id, to, subject, body string) Notice {
	return Notice{ID: id, Recipient: to, Subject: subject, Body: body}
}
func (n *Notice) Send(now time.Time) error {
	if n.Recipient == "" || n.Subject == "" {
		return fmt.Errorf("notice incomplete")
	}
	n.SentAt = now
	return nil
}
func (n *Notice) MarkRead()    { n.Read = true }
func (n Notice) Pending() bool { return !n.Read && n.SentAt.IsZero() == false }
