package handlers

import (
	"net/http"

	"github.com/caarlos0/httperr"
	"github.com/felipeweb/clean-arch/entity"
	"github.com/felipeweb/clean-arch/usecase"
	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func savePorts(service usecase.PortUsecase) http.Handler {
	return httperr.NewF(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		decoder := json.NewDecoder(r.Body)
		for decoder.More() {
			port := entity.Port{}
			err := decoder.Decode(&port)
			if err != nil {
				return httperr.Wrap(err, http.StatusBadRequest)
			}
			_, err = service.Save(ctx, port.Key, &port.PortInfo)
			if err != nil {
				return httperr.Wrap(err, http.StatusInternalServerError)
			}

		}
		defer r.Body.Close()
		w.WriteHeader(http.StatusAccepted)
		return nil
	})
}

func MakePortsHandler(r chi.Router, service usecase.PortUsecase) {
	r.Method(http.MethodPut, "/ports", savePorts(service))
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
}
