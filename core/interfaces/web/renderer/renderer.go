package renderer

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Renderer struct {
    templates *template.Template
}

func NewRenderer(componentsDir, pagesDir string) (*Renderer, error) {
    tmpl := template.New("")
    components, err := filepath.Glob(filepath.Join(componentsDir, "*.html"))
    if err != nil {
        return nil, err
    }
    if len(components) > 0 {
        tmpl, err = tmpl.ParseFiles(components...)
        if err != nil {
            return nil, err
        }
    }
    pages, err := filepath.Glob(filepath.Join(pagesDir, "*.html"))
    if err != nil {
        return nil, err
    }
    if len(pages) > 0 {
        tmpl, err = tmpl.ParseFiles(pages...)
        if err != nil {
            return nil, err
        }
    }
    return &Renderer{templates: tmpl}, nil
}

func (r *Renderer) Render(w http.ResponseWriter, tmplName string, data interface{}) error {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    return r.templates.ExecuteTemplate(w, tmplName, data)
}

