package handler

import (
	"html/template"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger *slog.Logger
	tmpl   *template.Template
}

func New(logger *slog.Logger, tmpl *template.Template) *Handler {
	return &Handler{ logger, tmpl }
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	h.tmpl.ExecuteTemplate(w, "index.html", nil)
}
