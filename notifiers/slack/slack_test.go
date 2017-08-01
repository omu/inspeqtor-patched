package slack

import (
	"strings"
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
	"github.com/uzem/inspeqtor-patched/utils"
)

func TestSlackNotifier(t *testing.T) {
	host := inspeqtor.NewHost()

	var T = map[string]string{
		"username":   "geek",
		"icon_emoji": "ghost",
		"url":        "https://acmecorp.slack.com/services/hooks/incoming-webhook?token=xxx/xxx/xxx",
	}

	action, err := BuildSlackNotifier(host, T)
	if err != nil {
		t.Fatalf("notifier 'slack' could not built: %v", err)
	}

	t.Parallel()
	util.LogInfo = true

	var theurl string
	var params message

	sendHere := func(url, _ string, v interface{}) error {
		theurl = url
		params = v.(message)
		return nil
	}

	sn := action.(*notifier)
	sn.senderFunc = sendHere

	alert := utils.MockEvent(host, inspeqtor.RuleFailed)
	if err := sn.Trigger(alert); err != nil {
		t.Errorf("notifier 'slack' could not execute: %v", err)
	}

	if T["url"] != theurl {
		t.Errorf("Expecting %s, got %s\n", T["url"], theurl)
	}

	expected := "localhost: swap is greater than 20%"
	if !strings.Contains(params.Text, expected) {
		t.Errorf("Expecting %s, got %s\n", expected, params.Text)
	}
}
