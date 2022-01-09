package main

import (
	"local.packages/bot"
	"local.packages/getContributions"
)

func main() {
	whurl := ""
	dw := &bot.DiscordWebhook{}
	content := bot.Content{}

	getContributions.GetContributions()

	if isContribution() {
		content.Text = "@channel 進捗どうですか？"
	}

	dw.CreateMessage(content)
	dw.SendWebhook(whurl)
}

func isContribution(num githubv4.int) bool {
	if num == 0 {
		return true
	}

	return false
}