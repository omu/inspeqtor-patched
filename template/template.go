package template

import (
	"io"
	"text/template"

	"github.com/mperham/inspeqtor"
)

const deftmpl = `
{{ if eq .Type.String "JobOverdue" }}{{.Hostname}}] Recurring job "{{.Thing.Name}}" is overdue.{{ end }}
{{ if eq .Type.String "JobRan" }}[{{.Hostname}}] Overdue job "{{.Thing.Name}}" just ran successfully.{{ end }}
{{ if eq .Type.String "ProcessDoesNotExist" }}{{.Hostname}}[{{.Thing.Name}}] does not exist.{{ end }}
{{ if eq .Type.String "ProcessExists" }}{{.Hostname}}[{{.Thing.Name}}] now running with PID {{.Service.Process.Pid}}{{ end }}
{{ if eq .Type.String "RuleFailed" }}{{.Target}}: {{.Rule.Metric}} is {{.Rule.Op}} {{.Rule.DisplayThreshold}}{{.Rule.Consequence}} {{ end }}
{{ if eq .Type.String "RuleRecovered" }}{{.Target}}: {{.Rule.Metric}} has recovered. {{ end }}
`

type Template struct {
	tmpl *template.Template
}

func New(event *inspeqtor.Event) *Template {
	return &Template{template.New(event.Name())}
}

func (t *Template) Parse() (*Template, error) {
	nt, err := t.tmpl.Clone()
	if err != nil {
		return nil, err
	}

	t.tmpl, err = nt.Parse(deftmpl)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Template) Execute(w io.Writer, data interface{}) error {
	return t.tmpl.Execute(w, data)
}
