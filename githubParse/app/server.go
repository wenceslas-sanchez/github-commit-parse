package app

import (
	"githubParse/githubParse/app/root"
	"githubParse/githubParse/app/users"
	"githubParse/githubParse/commit"
	"githubParse/githubParse/organization"
	"log"
	"net/http"
	"sync"
)

type ApplicationData struct {
	*organization.Organization
	*commit.NestedCounter
	USerCount *map[string]int
}

func RunApplication(data *ApplicationData) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", root.Handler(data.Organization))
	mux.Handle("/users/",
		&users.UserHandler{Organization: data.Organization,
			RWMutex: &sync.RWMutex{}})

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
