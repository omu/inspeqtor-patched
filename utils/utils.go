package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mperham/inspeqtor"
	"github.com/uzem/inspeqtor-patched/template"
)

func EventTemplate(event *inspeqtor.Event) (string, error) {
	tmpl, err := template.New(event).Parse()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, event); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func Sender(url, token string, params interface{}) error {
	b, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(b)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func MockEvent(check inspeqtor.Checkable, status inspeqtor.EventType) *inspeqtor.Event {
	return &inspeqtor.Event{
		Type:      status,
		Eventable: check,
		Rule: &inspeqtor.Rule{
			Entity:           check,
			MetricFamily:     "swap",
			MetricName:       "",
			Op:               inspeqtor.GT,
			DisplayThreshold: "20%",
			Threshold:        20,
			CurrentValue:     0,
			PerSec:           false,
			CycleCount:       1,
			TrippedCount:     0,
			State:            inspeqtor.Ok,
			Actions:          nil,
		},
	}
}
