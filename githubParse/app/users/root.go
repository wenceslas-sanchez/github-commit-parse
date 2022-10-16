package users

import (
	"html/template"
	"log"
	"net/http"
)

const usersRootTemplate string = `<h1>{{.Login}}</h1>
<table>
<tr style='text-align: left'>
	<th>Users</th>
</tr>

{{range $name, $member := .Members}}
<tr>
	<td><a href='./{{$name}}/'>{{$name}}</a></td>
</tr>

{{end}}
</table>
`

var reportRoot = template.Must(template.New("users-information").Parse(usersRootTemplate))

func (u *UserHandler) UsersRoot(w http.ResponseWriter) {
	if err := reportRoot.Execute(w, u.Organization); err != nil {
		log.Fatal(err)
	}
}
