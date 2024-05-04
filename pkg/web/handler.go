package web

import (
	"net/http"

	"github.com/Shikachuu/template-files/pkg"
)

type server struct {
    mux *http.ServeMux
    db  pkg.Database
}

func NewHTTPHandler(db pkg.Database) http.Handler {
    mux := http.NewServeMux()
    s := &server{mux: mux, db: db}
    s.addRoutes()
    return s.mux
}
