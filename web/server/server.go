package server

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func parseAllTemplates() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func registerRoutes() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
}

// StartServer register all routes and start listening to the server
func StartServer() {
	parseAllTemplates()
	registerRoutes()
	http.ListenAndServe(":8100", nil)
}
