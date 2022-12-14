package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

var Config *AppConfig

func SetConfig(a *AppConfig) {
	Config = a
}

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
