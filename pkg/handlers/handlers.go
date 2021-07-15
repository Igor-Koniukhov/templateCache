package handlers

import (
	"github.com/Igor-Koniukhov/cacheSample/pkg/config"
	"github.com/Igor-Koniukhov/cacheSample/pkg/models"
	"github.com/Igor-Koniukhov/cacheSample/pkg/render"
	"net/http"
)
//Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}
//NewRepo creates new repository
func NewRepo(a *config.AppConfig)*Repository  {
	return &Repository{
		App: a,
	}
}
//NewHandlers create new handlers, set from the main function(set the repository for the handlers).

func NewHandlers(r *Repository)  {
	Repo = r
}


func(m *Repository) HomePage(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["greeting"]="HowdyHo! This is HomePage!"
	render.RenderTemplate(w,r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func(m *Repository) AboutPage(w http.ResponseWriter, r *http.Request)  {
	stringMap := make(map[string]string)
	stringMap["hello"] = "Hello, again. This is about page"
	//send the data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
