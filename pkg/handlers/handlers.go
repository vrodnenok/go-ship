package handlers

import (
	"net/http"

	"github.com/vrodnenok/go-ship/pkg/config"
	"github.com/vrodnenok/go-ship/pkg/models"
	"github.com/vrodnenok/go-ship/pkg/render"
)

// Repo is the repository used by the nadlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers is the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handler for "home" page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About handler for "about" page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["author"] = "Rodnenok"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
