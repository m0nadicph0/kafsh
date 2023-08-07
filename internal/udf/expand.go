package udf

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/m0nadicph0/kafsh/internal/generators/seq"
	"math"
	"text/template"
)

var funcMap = template.FuncMap{
	"uuid": func() string {
		return uuid.NewString()
	},
	"seq": func() int {
		defer seq.Incr()
		if seq.Get() == math.MaxInt {
			seq.Reset()
		}
		return seq.Get()
	},
}

func Expand(expr string) (string, error) {
	buff := new(bytes.Buffer)
	tmpl, err := template.New("data").Funcs(funcMap).Parse(expr)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(buff, nil)
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}
