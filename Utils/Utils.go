package Utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"net/smtp"
	"os"
	"strconv"
	"time"
	"fmt"
)

func SendEmailToAdmin(subject string, message string) error {
	msg := []byte("To: sahrulramadhan684@gmail.com\r\n" + "Subject:" + subject + "\r\n" + "\r\n" + message + "\r\n")

	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
		"smtp.gmail.com")

	if err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("SMTP_EMAIL"),
		[]string{"sahrulramadhan684@gmail.com"},
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

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		dbClone := db.Session(&gorm.Session{})
		var total int64
		dbClone.Count(&total)

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

		ctx.Set("total_pages", totalPages)
		ctx.Set("page_size", pageSize)
		ctx.Set("page", page)
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}


func ConvertMonthToIndonesian(month string) (string, error) {
	// Parse the input string into a time.Time object
	t, err := time.Parse("2006-01", month)
	if err != nil {
		return "", fmt.Errorf("invalid month format: %v", err)
	}

	// Map English month names to Indonesian month names
	indonesianMonths := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	// Get the English month name
	englishMonth := t.Month().String()

	// Convert to Indonesian month name
	indonesianMonth, ok := indonesianMonths[englishMonth]
	if !ok {
		return "", fmt.Errorf("unknown month: %s", englishMonth)
	}

	// Return the formatted string
	return fmt.Sprintf("%s %d", indonesianMonth, t.Year()), nil
}

// ConvertLabelsToIndonesian converts a slice of "YYYY-MM" formatted strings to Indonesian month names.
func ConvertLabelsToIndonesian(labels []string) ([]string, error) {
	indonesianLabels := make([]string, len(labels))
	for i, label := range labels {
		indonesianLabel, err := ConvertMonthToIndonesian(label)
		if err != nil {
			return nil, err
		}
		indonesianLabels[i] = indonesianLabel
	}
	return indonesianLabels, nil
}