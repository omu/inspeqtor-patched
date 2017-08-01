package slack

import (
	"errors"

	"github.com/mperham/inspeqtor"
	"github.com/uzem/inspeqtor-patched/utils"
)

const (
	defaultUsername  = "Inspeqtor Patched"
	defaultIconEmoji = "ghost"
)

type message struct {
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Text      string `json:"text"`
}

type notifier struct {
	url       string
	username  string
	iconEmoji string

	senderFunc func(string, string, interface{}) error
}

func newNotifier(url, username, iconEmoji string) *notifier {
	return &notifier{
		url:        url,
		username:   username,
		iconEmoji:  iconEmoji,
		senderFunc: utils.Sender,
	}
}

func (n *notifier) Trigger(event *inspeqtor.Event) error {
	tmpl, err := utils.EventTemplate(event)
	if err != nil {
		return err
	}

	msg := message{
		Username:  n.username,
		IconEmoji: ":" + n.iconEmoji + ":",
		Text:      tmpl,
	}

	return n.senderFunc(n.url, "", msg)
}

func BuildSlackNotifier(_ inspeqtor.Eventable, config map[string]string) (inspeqtor.Action, error) {
	url := ""
	if config["url"] != "" {
		url = config["url"]
	} else {
		return nil, errors.New("notifier missing url")
	}

	username := defaultUsername
	if config["username"] != "" {
		username = config["username"]
	}

	iconEmoji := defaultIconEmoji
	if config["icon_emoji"] != "" {
		iconEmoji = config["icon_emoji"]
	}

	return newNotifier(url, username, iconEmoji), nil
}
