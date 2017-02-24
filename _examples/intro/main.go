package main

import (
	"fmt"
	"os"

	"github.com/biz/templates"
)

func main() {
	// Create a partial named base that will be used to wrap the contents of a view
	templates.AddPartial("base", `
		<!DOCTYPE HTML>
		<html>
			<body>
			{{ block "header" .}}
				<header>
					I am the header that will wrap your view
				</header>
			{{ end }}

			{{ block "contents" . }}{{end}}

			{{ block "footer" . }}
				<footer>
					I am the footer that will wrap your view
				</footer>
			{{ end }}
			</body>
		</html>
	`)

	// Create a view named home
	templates.AddView("home", `
		{{ define "contents" }}
			<p>
				Hello from the home view
			</p>
		{{ end }}
	`)

	// Create a view named about
	templates.AddView("about", `
		{{ define "header" }}
			About page with its own idea of what it wants in the header
		{{ end }}

		{{ define "contents" }}
			Hello from the about view
		{{ end }}
	`)

	// Parse the templates
	// NOTE: views and partiels should be created and parsed at startup.
	templates.Parse()

	fmt.Println("Render the home template with the base partial")
	templates.MustExecute(os.Stdout, "base", "home", nil)

	fmt.Println("\nRender the about template with the base partial")
	templates.MustExecute(os.Stdout, "base", "about", nil)
}
