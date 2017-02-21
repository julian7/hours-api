package handlers

import (
	"net/http"
	"strconv"

	"github.com/julian7/hours-api/models"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

// AllProjects client list handler
func (env *Env) AllProjects(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("projects.list")

	items, err := models.AllProjects(env.Conn, r.URL.Query())
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)
	err = jsonapiRuntime.MarshalManyPayload(w, items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetProject projects.show
func (env *Env) GetProject(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("projects.show")
	var vars = mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	item, err := models.FetchProject(env.Conn, id)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	err = jsonapiRuntime.MarshalOnePayload(w, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
