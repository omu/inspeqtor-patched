package twistapp

import (
	"errors"
	"strconv"

	"github.com/mperham/inspeqtor"
	"github.com/uzem/inspeqtor-patched/utils"
)

const threadAddAPIUrl = "https://api.twistapp.com/api/v2/threads/add"

type thread struct {
	ChannelId int    `json:"channel_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type notifier struct {
	channelID int
	token     string

	senderFunc func(string, string, interface{}) error
}

func newNotifier(token string, channelID int) *notifier {
	return &notifier{
		channelID:  channelID,
		token:      token,
		senderFunc: utils.Sender,
	}
}

func (n *notifier) Trigger(event *inspeqtor.Event) error {
	tmpl, err := utils.EventTemplate(event)
	if err != nil {
		return err
	}

	thrd := thread{
		ChannelId: n.channelID,
		Title:     tmpl,
		Content:   "",
	}

	return n.senderFunc(threadAddAPIUrl, n.token, thrd)
}

func BuildTwistappNotifier(_ inspeqtor.Eventable, config map[string]string) (inspeqtor.Action, error) {
	channelID := ""
	if config["channel_id"] != "" {
		channelID = config["channel_id"]
	} else {
		return nil, errors.New("notifier missing channel id")
	}

	token := ""
	if config["token"] != "" {
		token = config["token"]
	} else {
		return nil, errors.New("notifier missing token")
	}

	cidInt, err := strconv.Atoi(channelID)
	if err != nil {
		return nil, err
	}

	return newNotifier(token, cidInt), nil
}
