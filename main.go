package main

import (
	"os"
	"os/signal"
	"strconv"
	"time"
	"fmt"

	"local.packages/bot"
	"local.packages/getContributions"
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	tk := time.NewTicker(1 * time.Minute)
	defer tk.Stop()

	loop:
        for {
                select {
                case <- sc:
                    fmt.Println("stop bot")
                    break loop
                case <-tk.C:
                    if time.Now().Hour() == 20 {
						runBot()
					}
                }
        }
}

func runBot() {
	getContributions.EnvLoading()
	whurl := os.Getenv("WEBHOOK") //webhook URL
	fmt.Println(whurl)
	dw := &bot.DiscordWebhook{}
	content := bot.Content{}

	tc := getContributions.GetContributions()
	content.TodayContributions = strconv.Itoa(int(tc))
	if tc == 0 {
		content.Text = "@channel 進捗どうですか？"
	}

	dw.CreateMessage(content)
	dw.SendMessage(whurl)
}