package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ashfaqrobin/booking-app/pkg/config"
	"github.com/ashfaqrobin/booking-app/pkg/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

// Add default data to templatedata
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)

	return td
}

// Render a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tmplCache map[string]*template.Template
	var err error

	if config.Config.UseCache {
		tmplCache = config.Config.TemplateCache
	} else {
		tmplCache, err = CreateTemplateCache()

		if err != nil {
			log.Fatal(err)
		}
	}

	t, ok := tmplCache[tmpl]

	if !ok {
		log.Fatal("Template not found")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	buf.WriteTo(w)
}

// Create and return template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return tmplCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return tmplCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return tmplCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return tmplCache, err
			}
		}

		tmplCache[name] = ts
	}

	return tmplCache, nil

}
