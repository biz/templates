# Templates

The templates package provides a way to manage a collection of `html.Template`'s.


## Motivation

Go's `html/template` package is very powerful, it just needs a way to 
manage multiple templates, which the `templates` package was created to provide.

Prior to the introduction of the 'block' feature you could get away with chaining a bunch
of templates together and just use `Lookup` to retrieve a template and then call `ExecuteTemplate`
on said template. But if you are like me and you want to use `block` (lets you set up default values
and does not error out if you don't define a `block`) 
You quickly find out that you cannot define a block in multiple templates. [Playground Example](https://play.golang.org/p/6GBUT0-FyW)

## Basic Usage

The templates package separates the templates into two types, views and partials.

A partial is a template that will be made available to all defined views and other defined partials.

A view is a template that can use and execute against any defined partial. A view accesses partial, a partial does not
access a view

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
