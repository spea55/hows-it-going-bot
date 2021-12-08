package getContributions

import (
	"time"
	"github.com/google/go-github/v41/github"
	"github.com/google/go-github/github"
)

type Contributions struct {
	Date time.Time
	Count int
}

func GetContributions() {
	client := github.NewClient()
}