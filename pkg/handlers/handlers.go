package handlers

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bookings/pkg/render"
	"fmt"

	"net/http"
)

//repository used by handlers
var Repo *Repository

/*use repository pattern, it allows swap componants in application at a minimum
changes req to codebase
database connection can also share by using repository in future*/
type Repository struct {
	App *config.AppConfig
}

//creates a new repository
/*it takes pointer to appconfig and populate repository struct with it
return a new instance of it that application holds*/
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//provide access of repository to all handlers by creating method
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	/*working with session, grab remote IP address of person visiting
	my site and store it in home page*/
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	//pull remote IP out of session, string will be empty if there is no value
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	//send data to template
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}
