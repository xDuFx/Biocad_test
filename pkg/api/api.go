package api

import (
	"net/http"
	"testB/pkg/repository"

	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repository.PGRepo
}

func New(router *mux.Router, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/files/{page}/{limit}", api.files)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
