package handlers

import (
	"github.com/jbeausoleil/go-basic-web-application/pkg/config"
	"github.com/jbeausoleil/go-basic-web-application/pkg/render"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (*Repository) Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (*Repository) About(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
