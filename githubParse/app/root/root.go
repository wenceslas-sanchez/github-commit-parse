package root

import (
	"log"
	"net/http"

	"githubParse/githubParse/organization"
)

func Handler(o *organization.Organization) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		if err := Report.Execute(w, o); err != nil {
			log.Fatal(err)
		}
	}
}
