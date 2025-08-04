package web

import (
	"html/template"
	"net/http"

	"github.com/Fish-Barrel/drycharting/web/handler"
)

func (app *App) loadRoutes() {
	tmpl := template.Must(template.New("").ParseFS(app.templates, "templates/*"))
	handler := handler.New(app.logger, tmpl)
	files := http.FileServer(http.Dir("./static"))

	app.router.Handle("GET /static/", http.StripPrefix("/static", files))

	app.router.Handle("GET /", http.HandlerFunc(handler.Index))
}

