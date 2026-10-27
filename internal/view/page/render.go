//ff:func feature=page type=renderer control=sequence
//ff:what Render — write the page as one self-contained HTML document
package page

import (
	"embed"
	"html/template"
	"io"
)

//go:embed page.html
var assets embed.FS

// pageTemplate is the viewer page; it carries its own style and script.
var pageTemplate = template.Must(template.ParseFS(assets, "page.html"))

// Render writes p to w as one HTML document with inline style and script.
func Render(w io.Writer, p Page) error {
	return pageTemplate.Execute(w, p)
}
