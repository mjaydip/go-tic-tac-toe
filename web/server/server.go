package server

import (
	"html/template"
	"net/http"

	"github.com/mjaydip/go-examples/go-tic-tac-toe/web/game"
)

var funcMap template.FuncMap
var tpl *template.Template
var g game.Game

func parseAllTemplates() {
	funcMap = template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"mod": func(i int, m int) int {
			return i % m
		},
	}

	tpl = template.Must(template.New("main").Funcs(funcMap).ParseGlob("templates/*.html"))
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", g)
}

func registerRoutes() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))

	http.HandleFunc("/move", game.HandleMove)
	http.HandleFunc("/", index)
}

// StartServer register all routes and start listening to the server
func StartServer() {
	g.InitGame()
	parseAllTemplates()
	registerRoutes()
	http.ListenAndServe(":8100", nil)
}
