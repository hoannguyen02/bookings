package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/hoannguyen02/golang/bookings/pkg/config"
	"github.com/hoannguyen02/golang/bookings/pkg/models"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func GetRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func SetRepo(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "Home Page"
	stringMap["desc"] = "This is home page"
	renderPage(w, "home", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "About Page"
	stringMap["desc"] = "This is about page"
	renderPage(w, "about", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "generals", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "majors", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "make-reservation", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "contact", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "search-availability", &models.TemplateData{})
}

func renderPage(w http.ResponseWriter, page string, data *models.TemplateData) {
	tpl, err := template.ParseFiles("templates/base.tmpl", "templates/"+page+".page.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, data)
}
