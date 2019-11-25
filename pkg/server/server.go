package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benjigoldberg/fireplace/pkg/fireplace"
	"github.com/gorilla/mux"
)

// RegisterMuxes registers HTTP handlers with the webserver mux
func RegisterMuxes(mux *mux.Router) {
	mux.HandleFunc("/fireplace", fireplaceHandler)
}

func fireplaceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		state := fireplace.State{}
		if err := json.NewDecoder(r.Body).Decode(&state); err != nil {
			http.Error(w, fmt.Sprintf("Failed to unmarshal JSON: %v", err), http.StatusBadRequest)
			return
		}
		if err := state.Set(); err != nil {
			http.Error(w, fmt.Sprintf("Failed to set fireplace state: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
	}
}
