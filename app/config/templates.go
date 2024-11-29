package config

import (
	"html/template"
)

var Templates *template.Template

func InitTemplates(patterns ...string) error {
	Templates = template.New("")

	for _, pattern := range patterns {
		if pattern[0] == '/' {
			pattern = pattern[1:]
		}

		_, err := Templates.ParseGlob("./static/" + pattern)
		if err != nil {
			return err
		}
	}

	return nil
}
