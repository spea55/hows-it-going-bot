module github/spea/hows-it-going-bot

go 1.17

replace local.packages/bot => ./bot

replace local.packages/getContributions => ./getContributions

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-github/v41 v41.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/shurcooL/githubv4 v0.0.0-20220106005112-0707a5a90543 // indirect
	github.com/shurcooL/graphql v0.0.0-20200928012149-18c5c3165e3a // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20220105145211-5b0dc2dfae98 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	local.packages/bot v0.0.0-00010101000000-000000000000 // indirect
	local.packages/getContributions v0.0.0-00010101000000-000000000000 // indirect
)
