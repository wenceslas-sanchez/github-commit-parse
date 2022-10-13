package root

import (
	"log"
	"net/http"

	"githubParse/githubParse/organization"
)

func Handler(o *organization.Organization) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Report.Execute(w, o); err != nil {
			log.Fatal(err)
		}
	}
}
