package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ashfaqrobin/booking-app/internal/config"
	"github.com/ashfaqrobin/booking-app/internal/forms"
	"github.com/ashfaqrobin/booking-app/internal/models"
	"github.com/ashfaqrobin/booking-app/internal/render"
	"github.com/ashfaqrobin/booking-app/internal/structs"
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

// Contact page handler function
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
	var emptyReservation models.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// Handle posing of the reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 6)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		// TODO: Change render page from test to make-reservation
		render.RenderTemplate(w, r, "test.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	config.Config.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// Majors suite page handler function
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {

	reservation, ok := config.Config.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		config.Config.Session.Put(r.Context(), "flash", "You don't have any reservation")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	config.Config.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
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
	render.RenderTemplate(w, r, "test.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}
