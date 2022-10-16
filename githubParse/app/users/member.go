package users

import (
	"githubParse/githubParse/member"
	"html/template"
	"log"
	"net/http"
)

const memberTemplate string = `
<div>
	<img src='{{.User.AvatarURL}}' alt="logo" width='120' height='120'/>
	<h1><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a> (role {{.Role}})</h1>
</div>

`

var reportMember = template.Must(template.New("member-information").Parse(memberTemplate))

func MemberInformation(m *member.Member, w http.ResponseWriter) {
	if err := reportMember.Execute(w, *m); err != nil {
		log.Fatal(err)
	}
}
