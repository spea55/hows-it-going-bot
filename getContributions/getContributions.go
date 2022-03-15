package getContributions

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Contributions struct {
	Date time.Time
	Count int
}

func getTime() (time.Time, time.Time) {
	t := time.Now()
	y := time.Now().AddDate(0, 0, -1)
	
	return t, y
}

func GetContributions() (githubv4.Int){
	today, yesterday := getTime()

	err := godotenv.Load("getContributions/.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var t struct {
		User struct {
			ContrbutionsCollection struct {
				TotalCommitContributions githubv4.Int
			} `graphql:"contributionsCollection(from: $Yesterday,to: $Today)"`
		} `graphql:"user(login: \"spea55\")"`
	}

	variable := map[string]interface{}{
		"Today":     githubv4.DateTime{today},
		"Yesterday": githubv4.DateTime{yesterday},
	}

	err = client.Query(context.Background(), &t, variable)
	if err != nil {
		log.Fatalln(err)
	}
	tc := t.User.ContrbutionsCollection.TotalCommitContributions

	return tc
}