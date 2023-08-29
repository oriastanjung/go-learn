package handlers

import (
	"net/http"

	"github.com/oriastanjung/go-course/pkg/config"
	"github.com/oriastanjung/go-course/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// cara baru
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// render.RenderTemplate(w, "home.page.tmpl")
	// render.RenderTemplateMoreEficient(w, "home.page.tmpl")
	render.RenderTemplateMoreEficientAndComplex(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// render.RenderTemplate(w, "about.page.tmpl")
	// render.RenderTemplateMoreEficient(w, "about.page.tmpl")
	render.RenderTemplateMoreEficientAndComplex(w, "about.page.tmpl")
}

// // cara biasa untuk buat routes handler
// func Home(w http.ResponseWriter, r *http.Request) {
// 	// render.RenderTemplate(w, "home.page.tmpl")
// 	// render.RenderTemplateMoreEficient(w, "home.page.tmpl")
// 	render.RenderTemplateMoreEficientAndComplex(w, "home.page.tmpl")
// }

// func About(w http.ResponseWriter, r *http.Request) {
// 	// render.RenderTemplate(w, "about.page.tmpl")
// 	// render.RenderTemplateMoreEficient(w, "about.page.tmpl")
// 	render.RenderTemplateMoreEficientAndComplex(w, "about.page.tmpl")
// }
