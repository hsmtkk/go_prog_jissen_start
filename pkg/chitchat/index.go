package chitchat

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.Must(template.ParseFiles(files...)))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
