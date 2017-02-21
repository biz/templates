package templates

import (
	"html/template"
	"io"
)

// Default is used to contain the default templates instance
var Default = New()

func Parse() {
	Default.Parse()
}

func ParseDir(dir string, stripPrefix string) (*Templates, error) {
	return Default.ParseDir(dir, stripPrefix)
}

func AddView(name string, tmpl string) {
	Default.AddView(name, tmpl)
}

func AddPartial(name string, tmpl string) {
	Default.AddPartial(name, tmpl)
}

func AddFunc(name string, f interface{}) {
	Default.AddFunc(name, f)
}

func AddFuncs(funcMap template.FuncMap) {
	Default.AddFuncs(funcMap)
}

func Delims(left, right string) {
	Default.Delims(left, right)
}

func UseExts(extensions []string) {
	Default.UseExts(extensions)
}

func Render(baseView, view string, data interface{}) ([]byte, error) {
	return Default.Render(baseView, view, data)
}

func MustRender(baseView, view string, data interface{}) {
	Default.MustRender(baseView, view, data)
}

func RenderSingle(view string, data interface{}) ([]byte, error) {
	return Default.RenderSingle(view, data)
}

func MustRenderSingle(view string, data interface{}) {
	Default.MustRenderSingle(view, data)
}

func Execute(w io.Writer, baseView, view string, data interface{}) error {
	return Default.Execute(w, baseView, view, data)
}

func MustExecute(w io.Writer, baseView, view string, data interface{}) {
	Default.MustExecute(w, baseView, view, data)
}

func ExecuteSingle(w io.Writer, view string, data interface{}) error {
	return Default.ExecuteSingle(w, view, data)
}

func MustExecuteSingle(w io.Writer, view string, data interface{}) {
	Default.MustExecuteSingle(w, view, data)
}
