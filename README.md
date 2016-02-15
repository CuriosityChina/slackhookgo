SlackHookGo
=====================================
Golang Client for Slack Incomming Webhooks with Richly-Formatted Messages

Usage
-------------------------------------
``` go
package main

import (
	"log"

	"github.com/lowstz/slackhookgo"
)

func main() {
	var attachment slackhookgo.MessageAttachment
	attachment.Color = "danger"
	attachment.Pretext = "test pretext"
	attachment.Text = "test text"
	attachment.Title = "test title"
	attachment.TitleLink = "https://www.google.com/"
	msg := slackhookgo.NewSlackMessage("slack-bot", "test")
	msg.IconEmoji = ":ghost:"
	msg.AddAttachment(attachment)
	url := "https://hooks.slack.com/services/XXXXXXXXX/YYYYYYYYYYYYYYYYYYYY"
	err := slackhookgo.Send(url, msg)
	if err != nil {
		log.Println(err)
	}
}
```
