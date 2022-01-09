package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type discordImg struct {
	URL string `json:"url"`
	H   int    `json:"height"`
	W   int    `json:"width"`
}
type discordAuthor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon_url"`
}
type discordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type discordEmbed struct {
	Title  string         `json:"title"`
	Desc   string         `json:"description"`
	URL    string         `json:"url"`
	Color  int            `json:"color"`
	Image  discordImg     `json:"image"`
	Thum   discordImg     `json:"thumbnail"`
	Author discordAuthor  `json:"author"`
	Fields []discordField `json:"fields"`
}
 
type DiscordWebhook struct {
	UserName  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []discordEmbed `json:"embeds"`
	TTS       bool           `json:"tts"`
}

type Content struct {
	Text string 		`json:"text"`
	TodayContributions string `json:"todayContributions"`
	YesterdayContributions string `json:"yesterdayContributions"`
	Conpare string      `json:"conpare"`
	Continuous string 	`json:"continuous"`
}

func (dw *DiscordWebhook) CreateMessage(content Content) {
	dw.UserName = "Github Contribution"
	dw.Content = content.Text
	dw.Embeds = []discordEmbed{
		{
			Color: 0x550000,
			Fields: []discordField{
				{Name: "今日のContributions", Value: content.TodayContributions + "contributions", Inline: true},
				{Name: "昨日のContributions", Value: content.YesterdayContributions + "contributions", Inline: true},
			},
		},
		{
			Color: 1,
			Fields: []discordField{
				{Name: "前日比", Value: content.Conpare + "contributions", Inline: true},
				{Name: "継続日数", Value: content.Continuous + "日", Inline: true},
			},
		},
	}
}

func (dw *DiscordWebhook) SendWebhook(whurl string) {
	j, err := json.Marshal(dw)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
 
	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("new request err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
 
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	if resp.StatusCode == 204 {
		fmt.Println("sent", dw) //成功
	} else {
		fmt.Printf("%#v\n", resp) //失敗
	}
}