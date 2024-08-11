package blogrenderer

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"strings"

	goldmark "github.com/kaleocheng/goldmark"
)

//go:embed "templates/*"
var postTemplates embed.FS

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

func (p Post) SanitizedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}

func (p Post) BodyHTML() template.HTML {
	buf := bytes.Buffer{}
	if err := goldmark.Convert([]byte(p.Body), &buf); err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	if err := r.templ.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		return err
	}

	return nil
}
