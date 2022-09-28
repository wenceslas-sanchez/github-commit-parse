package utils

import (
	"net/http"
	"os"
)

const BaseURL = "https://api.github.com"

var GITHUBHEADER = http.Header{
	"Accept":        {"application/vnd.github+json"},
	"Authorization": {"Bearer " + os.Getenv("TEST_GITHUB_ACCESS_TOKEN")},
}
