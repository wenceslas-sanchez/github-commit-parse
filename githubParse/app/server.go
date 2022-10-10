package app

import (
	"githubParse/githubParse/app/root"
	"githubParse/githubParse/organization"
	"log"
	"net/http"
)

func RunApplication(organization *organization.Organization) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", root.Handler(organization))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
