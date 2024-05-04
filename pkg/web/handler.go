package web

import (
	"log/slog"
	"net/http"

	"github.com/Shikachuu/template-files/pkg"
)

type server struct {
	mux    *http.ServeMux
	db     pkg.Database
	logger *slog.Logger
}

func NewHTTPHandler(db pkg.Database, l *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	s := &server{mux: mux, db: db, logger: l}
	s.addRoutes()
	return s.mux
}
