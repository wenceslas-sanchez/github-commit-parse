package users

import (
	"fmt"
	"githubParse/githubParse/organization"
	"githubParse/githubParse/utils"
	"net/http"
	"regexp"
	"sync"
)

var (
	usersRootRegex = regexp.MustCompile(`^/users/?$`)
	getUserRegex   = regexp.MustCompile(`^/users/(.*?)/?$`)
)

type UserHandler struct {
	*organization.Organization
	*sync.RWMutex
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch {
	case usersRootRegex.MatchString(path):
		u.UsersRoot(w)
		return
	case getUserRegex.MatchString(path):
		result := getUserRegex.FindStringSubmatch(r.URL.Path)
		if len(result) < 2 {
			utils.NotFound(w)
			return
		}
		member, ok := u.Organization.Members[result[1]]
		fmt.Println(result)
		if !ok {
			utils.NotFound(w)
			return
		}
		MemberInformation(u.Organization, member, w)
		return
	default:
		utils.NotFound(w)
	}
}
