package handlers

import (
	"net/http"

	"github.com/ashfaqrobin/booking-app/pkg/config"
	"github.com/ashfaqrobin/booking-app/pkg/models"
	"github.com/ashfaqrobin/booking-app/pkg/render"
)

type Repository struct {
}

var Repo *Repository

func NewRepo() *Repository {
	return &Repository{}
}

func NewHandler(r *Repository) {
	Repo = r
}

// Home page handler function
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	config.Config.Session.Put(r.Context(), "IP", ip)

	stringMap := make(map[string]string)
	stringMap["demo"] = "Hello World!!!"

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About page handler function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	userIP := config.Config.Session.GetString(r.Context(), "IP")

	stringMap := make(map[string]string)
	stringMap["IP"] = userIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
