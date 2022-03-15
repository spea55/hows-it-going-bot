package main

import (
	"strconv"

	"local.packages/bot"
	"local.packages/getContributions"
)

func main() {
	whurl := "" //webhook URL
	dw := &bot.DiscordWebhook{}
	content := bot.Content{}

	tc := getContributions.GetContributions()
	content.TodayContributions = strconv.Itoa(int(tc))
	// if isContribution() {
	// 	content.Text = "@channel 進捗どうですか？"
	// }

	dw.CreateMessage(content)
	dw.SendMessage(whurl)
}

// func isContribution(num githubv4.int) bool {
// 	if num == 0 {
// 		return true
// 	}

// 	return false
// }