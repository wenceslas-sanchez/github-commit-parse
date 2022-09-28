package organization

import (
	. "githubParse/githubParse/member"
	. "githubParse/githubParse/repository"
)

type Organization struct {
	Login           string
	Name            string
	ID              int    `json:"id"`
	NodeID          string `json:"node_id"`
	Description     string
	Company         string
	Location        string
	Email           string
	Members         map[string]*Member
	MemberNames     []string `json:"member_names"`
	Repositories    map[string]*Repository
	RepositoryNames []string `json:"repository_names"`
}
