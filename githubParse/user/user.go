package user

import "encoding/json"

type User struct {
	Login            string
	ID               int    `json:"id"`
	NodeID           string `json:"node_id"`
	AvatarURL        string `json:"avatar_url"`
	GravatarURL      string `json:"gravatar_url"`
	URL              string `json:"url"`
	HTMLURL          string `json:"html_url"`
	FollowersURL     string `json:"followers_url"`
	FollowingURL     string `json:"following_url"`
	StarredURL       string `json:"starred_url"`
	SubscriptionsURL string `json:"subscriptions_url"`
	OrganizationsURL string `json:"organizations_url"`
	Name             string
	Company          string
	Location         string
	Email            string
	Followers        int
	Following        int
}

func (user *User) String() string {
	userJSON, err := json.MarshalIndent(*user, "", "")
	if err != nil {
		panic(err)
	}

	return string(userJSON)
}
