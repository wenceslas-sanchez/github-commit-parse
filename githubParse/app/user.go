package app

import (
	"fmt"
	"githubParse/githubParse/organization"
	"net/http"
	"regexp"

	"githubParse/githubParse/user"
)

var (
	userRootRegex = regexp.MustCompile("^/users$")
	getUserRegex  = regexp.MustCompile("^/users/.*$")
)

type userHandler struct {
	user.User
	organization.Organization
}

func (u *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	switch {
	case r.Method == http.MethodGet:
		fmt.Println()
		return
	default:
		notFound(w, r)
	}
}

func (u *userHandler) userRoot(w http.ResponseWriter) {

}
