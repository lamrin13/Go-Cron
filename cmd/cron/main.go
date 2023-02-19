package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/robfig/cron"
)

var wg sync.WaitGroup

func main() {
	c := cron.New()

	c.AddFunc("@every 5s", func() {
		file, err := os.OpenFile("temp.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("error while opening the file")
			os.Exit(1)
		}
		defer file.Close()
		if _, err := file.WriteString(time.Now().String() + "\n"); err != nil {
			fmt.Printf("error while writing the file")
			os.Exit(1)
		}
	})
	wg.Add(1)
	c.Start()

	wg.Wait()

}
