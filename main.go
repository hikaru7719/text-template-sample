package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	type TemplateSet struct {
		templateStr string
		data        interface{}
	}

	cases := []TemplateSet{
		TemplateSet{"{{ 1 }}", nil},
		TemplateSet{`hoge  {{- "fuga" -}}  hoge`, nil},
		TemplateSet{`{{ if "hoge" }} fuga {{ end }}`, nil},
		TemplateSet{`{{ range . }} {{.}}hoge {{ end }}`, []int{1, 2, 3, 4}},
		TemplateSet{`{{define "T1"}}ONE{{end }}{{template "T1" }}{{template "T1" }}`, nil},
		TemplateSet{`{{define "T1"}} {{.}} ONE {{end }}{{template "T1" "ZERO"}}`, nil},
		TemplateSet{`{{block "T1" "ZERO"}} {{.}} ONE {{end }}`, nil},
		TemplateSet{`{{with "ZERO"}} {{.}} ONE {{end}}`, nil},
		TemplateSet{`{{$a := "hoge"}}{{$a}}`, nil},
		TemplateSet{`{{range $key, $value := .}}key:{{$key}},value:{{$value}}
{{end}}`, map[string]string{"hoge": "fuga", "foo": "bar"}},
	}

	for _, t := range cases {
		tmpl, err := template.New("test").Parse(t.templateStr)

		if err != nil {
			panic(err)
		}

		fmt.Println("--------")
		fmt.Printf("template sring is '%s'\n", t.templateStr)
		err = tmpl.Execute(os.Stdout, t.data)
		fmt.Println("\n--------")
		if err != nil {
			panic(err)
		}
	}
}
