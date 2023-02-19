package main

import (
	"fmt"
	"os"

	"github.com/lamrin13/cron/package/api"
)

func main() {
	mailContent, err := api.ParseJSON()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	api.SendEmail(mailContent)
}
