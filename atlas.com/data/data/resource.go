package data

import (
	"atlas-data/rest"
	"github.com/Chronicle20/atlas-rest/server"
	"github.com/Chronicle20/atlas-tenant"
	"github.com/gorilla/mux"
	"github.com/jtumidanski/api2go/jsonapi"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InitResource(si jsonapi.ServerInformation) server.RouteInitializer {
	return func(router *mux.Router, l logrus.FieldLogger) {
		r := router.PathPrefix("/data").Subrouter()
		r.HandleFunc("", rest.RegisterHandler(l)(si)("upload", uploadData)).Methods(http.MethodPatch)
	}
}

func uploadData(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := tenant.MustFromContext(d.Context())
		d.Logger().Debugf("Processing .zip for tenant [%s], region [%s], version [%d.%d].", t.Id().String(), t.Region(), t.MajorVersion(), t.MinorVersion())
		r.Body = http.MaxBytesReader(w, r.Body, 1<<30)

		err := r.ParseMultipartForm(1 << 30) // 1GB max file size
		if err != nil {
			d.Logger().WithError(err).Errorf("Unable to process zip.")
			w.WriteHeader(http.StatusRequestHeaderFieldsTooLarge)
			return
		}

		// Get file from request
		file, handler, err := r.FormFile("zip_file")
		if err != nil {
			d.Logger().WithError(err).Errorf("Unable to process zip.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		err = ProcessZip(d.Logger())(d.Context())(file, handler)
		if err != nil {
			d.Logger().WithError(err).Errorf("Unable to process zip.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
