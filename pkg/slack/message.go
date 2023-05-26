package slack

import (
	"fmt"
	"github.com/slack-go/slack"
	"time"
)

func DefaultAttachment(title, process, text, priority string) slack.Attachment {
	return slack.Attachment{
		Color:      "#0000FF",
		AuthorName: "ff.primrose",
		Title:      fmt.Sprintf("finexblock-matching-engine error : %v", title),
		Pretext:    fmt.Sprintf("ERROR IN PROCESS : %v", process),
		Text:       text,
		Fields: []slack.AttachmentField{
			{
				Title: "Priority",
				Value: priority,
				Short: false,
			},
			{
				Title: "Due Date",
				Value: time.Now().String(),
				Short: true,
			},
		},
	}
}
