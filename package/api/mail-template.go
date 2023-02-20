package api

import (
	"bytes"
	"fmt"
	"html/template"
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
