package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ashfaqrobin/booking-app/pkg/config"
	"github.com/ashfaqrobin/booking-app/pkg/models"
	"github.com/ashfaqrobin/booking-app/pkg/render"
	"github.com/ashfaqrobin/booking-app/pkg/structs"
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

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About page handler function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	userIP := config.Config.Session.GetString(r.Context(), "IP")

	stringMap := make(map[string]string)
	stringMap["IP"] = userIP

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals quaters page handler function
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

// Generals quaters page handler function
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

// Majors suite page handler function
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

// Make reservation page handler function
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

// Search availability page handler function
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

// Post search availability page handler function
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	fmt.Println(start, end)

	w.Write([]byte("Posted to search availability"))
}

// Check for search availability and send JSON
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	res := structs.JSONResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(res, "", "    ")

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(out)
}

// Search availability page handler function
func (m *Repository) Test(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "test.page.html", &models.TemplateData{})
}
