package handlers

import (
	"net/http"
	"strconv"

	"github.com/julian7/hours-api/models"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

// AllClients client list handler
func (env *Env) AllClients(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("clients.list")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)

	items, err := models.AllClients(env.Conn)
	if err != nil {
		return
	}
	err = jsonapiRuntime.MarshalManyPayload(w, items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetClient clients.show
func (env *Env) GetClient(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("clients.show")
	var vars = mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	item, err := models.FetchClient(env.Conn, id)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)
	err = jsonapiRuntime.MarshalOnePayload(w, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
