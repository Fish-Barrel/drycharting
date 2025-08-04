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

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	h.tmpl.ExecuteTemplate(w, "login.html", nil)
}

func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	h.tmpl.ExecuteTemplate(w, "welcome.html", nil)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}
