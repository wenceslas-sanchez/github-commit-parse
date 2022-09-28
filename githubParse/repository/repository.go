package repository

import (
	. "githubParse/githubParse/user"
)

const RepositoryURL = "/repos"

type Repository struct {
	Name         string
	FullName     string `json:"full_name"`
	ID           int    `json:"id"`
	NodeID       string `json:"node_id"`
	URL          string `json:"url"`
	HTMLURL      string `json:"html_url"`
	Description  string
	Owner        *User
	Contributors []string
}
