package templates_test

import (
	"os"

	"github.com/biz/templates"
)

// Basic usage using the templates Default templates.
//
// Below one partial named 'base' will be used to wrap two views, 'home' and 'about'
func Example() {
	// Create a partial named base that will be used to wrap the contents of a view
	templates.AddPartial("base", `
<!DOCTYPE HTML>
<html>
	<body>
	{{- block "header" .}}
		<header>
			I am the header that will wrap your view
		</header>
	{{- end }}

	{{- block "contents" . }}{{- end}}

	{{- block "footer" . }}
		<footer>
			I am the footer that will wrap your view
		</footer>
	{{- end }}
	</body>
</html>`)

	// Create a view named home
	templates.AddView("home", `
{{- define "contents" }}
	<p>
		Hello from the home view
	</p>
{{- end }}
	`)

	// Create a view named about
	templates.AddView("about", `
{{- define "header" }}
	<h1>About page with its own idea of what it wants in the header</h1>
{{- end }}

{{- define "contents" }}
	Hello from the about view
{{- end }}
	`)

	// Parse the templates
	// NOTE: views and partiels should be created and parsed at startup.
	templates.Parse()

	// Render the home template wrapped by the base partial
	templates.MustExecute(os.Stdout, "base", "home", nil)

	// Render the about template wrapped by the base partial
	templates.MustExecute(os.Stdout, "base", "about", nil)
	// Output:
	// <!DOCTYPE HTML>
	// <html>
	// 	<body>
	// 		<header>
	// 			I am the header that will wrap your view
	// 		</header>
	// 	<p>
	// 		Hello from the home view
	// 	</p>
	// 		<footer>
	// 			I am the footer that will wrap your view
	// 		</footer>
	// 	</body>
	// </html>
	// <!DOCTYPE HTML>
	// <html>
	// 	<body>
	// 	<h1>About page with its own idea of what it wants in the header</h1>
	// 	Hello from the about view
	// 		<footer>
	// 			I am the footer that will wrap your view
	// 		</footer>
	// 	</body>
	// </html>
}
