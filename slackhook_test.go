package slackhookgo

import "testing"

func TestNewSlackMessage(t *testing.T) {
	msg := NewSlackMessage()
	if len(msg.Attachments) != 0 {
		t.Errorf("want %d, got %d", 0, len(msg.Attachments))
	}

	msg = NewSlackMessage(MessageAttachment{})
	if len(msg.Attachments) != 1 {
		t.Errorf("want %d got %d", 1, len(msg.Attachments))
	}

	msg = NewSlackMessage(MessageAttachment{}, MessageAttachment{})
	if len(msg.Attachments) != 2 {
		t.Errorf("want %d got %d", 2, len(msg.Attachments))
	}
}
