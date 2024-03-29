package Mail

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"

	ST "github.com/YahirKeym/transactions/src/internal/transaction/summarize"
)

type Email struct {
	To string
}

func Send(summaryTransaction ST.SummaryTransaction, email Email) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	template, err := loadTemplate("src/internal/mail/template.html")
	if err != nil {
		return err
	}
	body, err := executeTemplate(template, summaryTransaction)
	if err != nil {
		return err
	}
	to := email.To
	message := buildEmailMessage(from, to, "STORI - Account Status", body)

	err = smtp.SendMail(host+":"+port,
		smtp.PlainAuth("", from, password, host),
		from, []string{to}, []byte(message))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	return nil
}

func loadTemplate(templatePath string) (*template.Template, error) {
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	return template, nil
}

func executeTemplate(template *template.Template, summaryTransaction ST.SummaryTransaction) (string, error) {
	var bodyBuffer bytes.Buffer
	err := template.Execute(&bodyBuffer, summaryTransaction)
	if err != nil {
		return "", err
	}
	return bodyBuffer.String(), nil
}

func buildEmailMessage(from, to, subject, body string) []byte {
	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		body + "\r\n"

	return []byte(message)
}
