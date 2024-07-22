package Utils

import (
	"github.com/joho/godotenv"
	"net/smtp"
	"os"
)

func SendEmailToAdmin(subject string, message string) error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	msg := []byte("To: mahendrakrs448@gmail.com\r\n" + "Subject:" + subject + "\r\n" + "\r\n" + message + "\r\n")

	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
		"smtp.gmail.com")

	if err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("SMTP_EMAIL"),
		[]string{"mahendrakrs448@gmail.com"},
		msg,
	); err != nil {
		return err
	}

	return nil
}

func SendEmail(to string, subject string, message string) error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	msg := []byte("To:" + to + "\r\n" + "Subject:" + subject + "\r\n" + "\r\n" + message + "\r\n")

	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
		"smtp.gmail.com")

	if err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("SMTP_EMAIL"),
		[]string{to},
		msg,
	); err != nil {
		return err
	}

	return nil
}
