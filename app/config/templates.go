package config

import (
	"html/template"
)

var Templates *template.Template

func InitTemplates() error {
	var err error
	Templates, err = template.ParseGlob("./static/pages/*html")
	return err
}
