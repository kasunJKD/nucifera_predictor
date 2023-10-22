package auth

import (
	"database/sql"
	"log"
	"net/smtp"
	"nucifera_backend/internal/db"
	"time"

	"github.com/go-co-op/gocron"
)

//not implemented just the base code for emails
//need to test and try implementing this
func sendEmail(user db.User, emailContent string) error {
	from := "kasun.ravihara@gmail.com"
	password := "xxx-xxx"
	to := user.Email
		
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(emailContent))

	if err != nil {
		return err
	}

	return nil
}

func sendEmailFull (database *sql.DB) interface{}{
	cc := db.DBConfig{
		DB: database,
	}

	users, err := cc.GetEmails()

	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		emailcontent := generatedEmailContent(user)
		err := sendEmail(user, emailcontent)
		if err != nil {
			log.Fatal("error sending email")
		}
	}

	select {}

}

func ScheduleEmails(database *sql.DB) {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(1).Month().At("00:00").Do(sendEmailFull(database))
	if err != nil {
		log.Fatal(err)
	}
	s.StartAsync()
}

func generatedEmailContent(user db.User) string {
	subject := "Hello, "+user.Email
	body := "Dear "+user.Email+ ",\n\n"

	emailContent := "subject " + subject
	emailContent += body

	return emailContent


}


