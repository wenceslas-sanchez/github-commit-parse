package member

import (
	"encoding/json"
	. "githubParse/githubParse/user"
	"time"
)

type Member struct {
	*User
	Inviter   struct{ User }
	Role      string
	CreatedAt time.Time `json:"created_at"`
	FailedAt  time.Time `json:"failed_at"`
}

func (member *Member) String() string {
	userJSON, err := json.MarshalIndent(*member, "", "")
	if err != nil {
		panic(err)
	}
	return string(userJSON)
}
