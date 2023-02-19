package api

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func MailBody() string {
	t, err := template.ParseFiles("C:\\Users\\Nirmal\\Desktop\\Go tutorials\\Cron\\package\\templates\\template.html")
	if err != nil {
		fmt.Println("Error while parsing template", err)
		os.Exit(1)
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, nil); err != nil {
		fmt.Println("Error while parsing template", err)
		os.Exit(1)
	}
	return buf.String()
}

func SendEmail(msg string) {
	from := "patelnirmal13595@gmail.com"
	password := "xlciuczmmfpwmxmw"

	toList := []string{"njpatel13595@gmail.com", "rutvip98@gmail.com"}

	host := "smtp.gmail.com"

	port := "587"

	// msg := api.MailBody()

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Test Email" + "!\n"

	body := []byte(subject + mime + "\n" + msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to all user in toList")

}
