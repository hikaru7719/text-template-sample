# text-template-sample

template sample

```
--------
template sring is '{{ 1 }}'
1
--------
--------
template sring is 'hoge  {{- "fuga" -}}  hoge'
hogefugahoge
--------
--------
template sring is '{{ if "hoge" }} fuga {{ end }}'
 fuga
--------
--------
template sring is '{{ range . }} {{.}}hoge {{ end }}'
 1hoge  2hoge  3hoge  4hoge
--------
--------
template sring is '{{define "T1"}}ONE{{end }}{{template "T1" }}{{template "T1" }}'
ONEONE
--------
--------
template sring is '{{define "T1"}} {{.}} ONE {{end }}{{template "T1" "ZERO"}}'
 ZERO ONE
--------
--------
template sring is '{{block "T1" "ZERO"}} {{.}} ONE {{end }}'
 ZERO ONE
--------
--------
template sring is '{{with "ZERO"}} {{.}} ONE {{end}}'
 ZERO ONE
--------
--------
template sring is '{{$a := "hoge"}}{{$a}}'
hoge
--------
--------
template sring is '{{range $key, $value := .}}key:{{$key}},value:{{$value}}
{{end}}'
key:foo,value:bar
key:hoge,value:fuga

--------
```
