package twistapp

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
	"github.com/uzem/inspeqtor-patched/utils"
)

func TestTwistappNotifier(t *testing.T) {
	host := inspeqtor.NewHost()

	var T = map[string]string{
		"channel_id": "43567",
		"token":      "oauth2:d25cb231353c9fff6ilb43ef3a907f1533fb5a05",
	}

	action, err := BuildTwistappNotifier(host, T)
	if err != nil {
		t.Fatalf("notifier 'twistapp' could not built: %v", err)
	}

	t.Parallel()
	util.LogInfo = true

	var token string
	var params thread

	sendHere := func(_, tok string, v interface{}) error {
		token = tok
		params = v.(thread)
		return nil
	}

	sn := action.(*notifier)
	sn.senderFunc = sendHere

	alert := utils.MockEvent(host, inspeqtor.RuleFailed)
	if err := sn.Trigger(alert); err != nil {
		t.Errorf("notifier 'twistapp' could not execute: %v", err)
	}

	if T["token"] != token {
		t.Errorf("Expecting %s, got %s\n", T["token"], token)
	}

	cidInt, err := strconv.Atoi(T["channel_id"])
	if err != nil {
		t.Errorf("Internal error: %v", err)
	}

	if params.ChannelId != cidInt {
		t.Errorf("Expecting %s, got %s\n", T["channel_id"], params.ChannelId)
	}

	expectedTitle := "localhost: swap is greater than 20%"
	if !strings.Contains(params.Title, expectedTitle) {
		t.Errorf("Expecting %s, got %s\n", expectedTitle, params.Title)
	}
}
