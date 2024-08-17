package Utils

import (
	"net/smtp"
	"os"
)

func SendEmailToAdmin(subject string, message string) error {
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

func MonthToRoman(month int) string {
	months := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
	}

	roman, exists := months[month]
	if !exists {
		return ""
	}

	return roman
}
