package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/biz/templates"
)

var (
	// templates global that will contain all of our parsed temlates from the templates directory
	tmpls *templates.Templates
	port  = ":8083"
	host  = "http://localhost" + port
)

var (
	css     = []string{fmt.Sprintf("%v/static/css/main.css", host)}
	scripts = []string{fmt.Sprintf("%v/static/js/main.js", host)}
)

// parse the templates in the template directory
func init() {
	var err error
	tmpls, err = templates.New().ParseDir("./templates", "templates/")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tmpls.Parse()

	// web pages
	http.HandleFunc("/", renderPage("index", "Index Page Title"))
	http.HandleFunc("/about", renderPage("about", "About Page Title"))

	// serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start http server
	log.Println("Server started on " + host)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func renderPage(view string, title string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := tmpls.RenderTemplate("base.html", "views/"+view+".html", map[string]interface{}{
			"Title":   title,
			"Css":     css,
			"Scripts": scripts,
			"Menu":    activeNav(view),
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(b)
	}
}

type navItem struct {
	Name  string
	Attrs map[template.HTMLAttr]string
}

func activeNav(active string) []navItem {
	// create menu items
	about := navItem{
		Name: "About",
		Attrs: map[template.HTMLAttr]string{
			"href":  "/about",
			"title": "About Page",
		},
	}
	home := navItem{
		Name: "Home",
		Attrs: map[template.HTMLAttr]string{
			"href":  "/",
			"title": "Home Page",
		},
	}

	// set active menu class
	switch active {
	case "about":
		about.Attrs["class"] = "active"
	case "home":
		home.Attrs["class"] = "active"
	}

	return []navItem{home, about}
}
