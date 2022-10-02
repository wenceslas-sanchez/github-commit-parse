package organization

import "html/template"

const Template string = `<h1>{{.Login}}</h1>
<table>
<tr style='text-align: left'>
	<th>Repository</th>
	<th>Description</th>
	<th>Link</th>
</tr>

{{range $name, $repository := .Repositories}}
<tr>
	<td><a href='./{{$name}}'>{{$name}}</a></td>
	<td>{{$repository.Description}}</td>
	<td><a href='{{$repository.HTMLURL}}'>Github</a></td>
</tr>

{{end}}
</table>`

var Report = template.Must(template.New("organization-information").Parse(Template))
