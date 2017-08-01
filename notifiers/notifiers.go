package notifiers

import (
	"github.com/mperham/inspeqtor"
	"github.com/uzem/inspeqtor-patched/notifiers/slack"
	"github.com/uzem/inspeqtor-patched/notifiers/twistapp"
)

var notifiers = map[string]inspeqtor.NotifierBuilder{
	"slack":    slack.BuildSlackNotifier,
	"twistapp": twistapp.BuildTwistappNotifier,
}

func init() {
	for name, builder := range notifiers {
		inspeqtor.Notifiers[name] = builder
	}
}
