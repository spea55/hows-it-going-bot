module github/spea/hows-it-going-bot

go 1.17

replace local.packages/bot => ./bot

replace local.packages/getContributions => ./getContributions

require (
	github.com/google/go-github/v41 v41.0.0 // indirect
	local.packages/bot v0.0.0-00010101000000-000000000000 // indirect
	local.packages/getContributions v0.0.0-00010101000000-000000000000 // indirect
)
