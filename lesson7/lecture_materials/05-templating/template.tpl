Grade: {{ .Grade }}
Teacher: {{ .Teacher.FirstName }} {{ .Teacher.LastName }}, {{ .Teacher.Age }}
Students:
{{ range $key, $val := .Students }}  {{ inc $key }}: {{ $val.FirstName }} {{ $val.LastName }}, {{ $val.Age | printf "\"%d\"" }}
{{ end }}
