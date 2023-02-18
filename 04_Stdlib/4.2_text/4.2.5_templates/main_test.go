package main

import (
	"testing"
	"text/template"
)

func Test(t *testing.T) {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))

	user := User{"Алиса", 500}
	got := renderToString(tpl, user)

	const want = "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке."
	if got != want {
		t.Errorf("%v: got '%v'", user, got)
	}
}
