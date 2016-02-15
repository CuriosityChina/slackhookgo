package slackhookgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackMessage struct {
	Username    string              `json:"username,omitempty"`
	IconUrl     string              `json:"icon_url,omitempty"`
	IconEmoji   string              `json:"icon_emoji,omitempty"`
	Channel     string              `json:"channel,omitempty"`
	Text        string              `json:"text,omitempty"`
	Attachments []MessageAttachment `json:"attachments,omitempty"`
}

type MessageAttachment struct {
	Fallback   string            `json:"fallback,omitempty"`
	Color      string            `json:"color,omitemtpy"`
	Pretext    string            `json:"pretext,omitempty"`
	AuthorName string            `json:"author_name,omitemtpy"`
	AuthorLink string            `json:"author_link,omitemtpy"`
	AuthorIcon string            `json:"author_icon,omitempty"`
	Title      string            `json:"title,omitempty"`
	TitleLink  string            `json:"title_link,omitempty"`
	Text       string            `json:"text,omitempty"`
	Fields     []AttachmentField `json:"fields,omitempty"`
	ImageURL   string            `json:"image_url,omitempty"`
	ThumbURL   string            `json:"thumb_url,omitempty"`
}

type AttachmentField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

func NewSlackMessage(username, channel string) *SlackMessage {
	var msg SlackMessage
	msg.Username = username
	msg.Channel = channel
	return &msg
}

func (m *SlackMessage) AddAttachment(attachment MessageAttachment) *SlackMessage {
	m.Attachments = append(m.Attachments, attachment)
	return m
}

func (m *SlackMessage) AddAttachments(attachments []MessageAttachment) *SlackMessage {
	m.Attachments = append(m.Attachments, attachments...)
	return m
}

func Send(url string, msg *SlackMessage) error {
	buf, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
