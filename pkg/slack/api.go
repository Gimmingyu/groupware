package slack

import "github.com/slack-go/slack"

func NewApiClient(token string) *slack.Client {
	return slack.New(token)
}

func PostMessage(api *slack.Client, channelID string, message string, attachments ...slack.Attachment) (string, string, error) {
	return api.PostMessage(
		channelID,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAttachments(attachments...),
		slack.MsgOptionAsUser(false),
	)
}
