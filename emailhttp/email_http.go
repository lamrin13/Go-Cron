package emailhttp

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lamrin13/cron/package/api"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("SendEmail", SendEmail)
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	mailContent, err := api.ParseJSON()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	api.SendEmail(mailContent)

	fmt.Fprint(w, "Sent eamil successfully")
}
