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
