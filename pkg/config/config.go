package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig creates a global config struct
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	PORT_NUMBER   string
	InProduction  bool
	Session       *scs.SessionManager
}
