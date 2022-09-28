package commit

import "githubParse/githubParse/user"

type Commit struct {
	URL       string `json:"url"`
	Author    *user.User
	Committer *user.User
	Commit    struct {
		message string
	}
}
