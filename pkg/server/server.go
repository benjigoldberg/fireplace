package server

import (
	"fmt"
	"net/http"

	"github.com/benjigoldberg/fireplace/pkg/fireplace"
	"github.com/gorilla/mux"
	"github.com/spothero/tools/log"
)

// RegisterMuxes registers HTTP handlers with the webserver mux
func RegisterMuxes(mux *mux.Router) {
	mux.HandleFunc("/fireplace", fireplaceHandler)
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
}

func fireplaceHandler(w http.ResponseWriter, r *http.Request) {
	logger := log.Get(r.Context())
	switch r.Method {
	case http.MethodPost:
		logger.Info("Received POST")
		if err := r.ParseForm(); err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse form: %v", err), http.StatusBadRequest)
			return
		}
		state := fireplace.State{}
		fireplaceState := r.PostFormValue("fireplaceState")
		switch fireplaceState {
		case "allOn":
			state.Flame = true
			state.BlowerFan = true
		case "fireplaceOn":
			state.Flame = true
			state.BlowerFan = false
		case "allOff":
			state.Flame = false
			state.BlowerFan = false
		default:
			http.Error(w, fmt.Sprintf("invalid fireplace state received: %s", fireplaceState), http.StatusBadRequest)
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
