package main

import (
	"bytes"
	"text/template"
)

// начало решения

var templateText = `{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. {{if ge .Balance 100}}Все в порядке.{{else if gt .Balance 0}}Пора пополнить.{{else}}Доступ заблокирован.{{end}}`

// конец решения

type User struct {
	Name    string
	Balance int
}

// renderToString рендерит данные по шаблону в строку
func renderToString(tpl *template.Template, data any) string {
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	return buf.String()
}
