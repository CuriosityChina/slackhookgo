package slackhookgo

import "testing"

func TestNewSlackMessage(t *testing.T) {
	msg := NewSlackMessage("test-bot", "test")

	if len(msg.Attachments) != 0 {
		t.Errorf("want %d, got %d", 0, len(msg.Attachments))
	}

	msg.AddAttachment(MessageAttachment{})
	if len(msg.Attachments) != 1 {
		t.Errorf("want %d got %d", 1, len(msg.Attachments))
	}

	msg.AddAttachment(MessageAttachment{})
	if len(msg.Attachments) != 2 {
		t.Errorf("want %d got %d", 2, len(msg.Attachments))
	}
}
