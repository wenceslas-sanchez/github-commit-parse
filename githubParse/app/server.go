package app

import (
	"githubParse/githubParse/organization"
	"log"
	"net/http"
)

func RunApplication(organization *organization.Organization) {
	http.HandleFunc("/", handler(organization))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(o *organization.Organization) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := organization.Report.Execute(w, o); err != nil {
			log.Fatal(err)
		}
	}
}
