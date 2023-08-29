package config

import "html/template"

//AppConfig holds the application config globally

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}
