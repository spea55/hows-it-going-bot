package main

import (
	"local.packages/bot"
	"local.packages/getContributions"
)

func main() {
	var content bot.Content
	var contribution int

	if isContribution(contribution) {
		content.Text = "@channel 進捗どうですか？"
	}

	bot.Webhook(content)
}

func isContribution(num int) bool {
	if num == 0 {
		return true
	}

	return false
}