package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/resend/resend-go/v2"
)

//go:embed static/*.html
var staticFS embed.FS

type Mailer struct {
	ResendApiKey string
	tpl          *template.Template
}

// NewMailer builds the mailer and pre-loads every /static/*.html file.
func NewMailer(apiKey string) (*Mailer, error) {
	tpl := template.New("")

	files, err := staticFS.ReadDir("static")

	if err != nil {
		return nil, fmt.Errorf("read static dir: %w", err)
	}

	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".html" {
			continue
		}

		content, err := staticFS.ReadFile("static/" + f.Name())

		if err != nil {
			return nil, fmt.Errorf("read %s: %w", f.Name(), err)
		}

		name := f.Name()[:len(f.Name())-5] // remove .html

		tpl, err = tpl.New(name).Parse(string(content))

		if err != nil {
			return nil, fmt.Errorf("parse %s: %w", f.Name(), err)
		}
	}

	return &Mailer{
		ResendApiKey: apiKey,
		tpl:          tpl,
	}, nil
}

// SendTemplate renders the named template with data and sends the email.
func (m *Mailer) SendTemplate(name string, data any, to, subject string) (*resend.SendEmailResponse, error) {
	var buf bytes.Buffer
	if err := m.tpl.ExecuteTemplate(&buf, name, data); err != nil {
		return nil, fmt.Errorf("execute template %s: %w", name, err)
	}
	return m.Send(buf.String(), to, subject)
}

func (m *Mailer) Send(body, to, subject string) (*resend.SendEmailResponse, error) {
	// initialize client
	client := resend.NewClient(m.ResendApiKey)

	params := &resend.SendEmailRequest{
		To:      []string{to},
		From:    "WatchCircle <onboarding@resend.dev>",
		Html:    body,
		Subject: subject,
	}

	return client.Emails.Send(params)
}
