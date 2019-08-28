package notifiers

import (
    "fmt"

    "github.com/nlopes/slack"
)

func SendMessageToChannel(channelID string, message string)  {
	api := slack.New("YOUR_TOKEN_HERE")
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text: message,
	}

	id, timestamp, err := api.PostMessage(channelID, slack.MsgOptionText("WNCP First Contact", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", id, timestamp)
}