package equipment

import (
	"atlas-data/rest"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/Chronicle20/atlas-rest/server"
	"github.com/gorilla/mux"
	"github.com/jtumidanski/api2go/jsonapi"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InitResource(si jsonapi.ServerInformation) server.RouteInitializer {
	return func(router *mux.Router, l logrus.FieldLogger) {
		registerGet := rest.RegisterHandler(l)(si)

		r := router.PathPrefix("/data/equipment").Subrouter()
		r.HandleFunc("/{equipmentId}", registerGet("get_equipment_statistics", handleGetEquipmentStatistics)).Methods(http.MethodGet)
		r.HandleFunc("/{equipmentId}/slots", registerGet("get_equipment_slots", handleGetEquipmentSlots)).Methods(http.MethodGet)
	}
}

func handleGetEquipmentStatistics(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseEquipmentId(d.Logger(), func(equipmentId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			e, err := GetById(d.Context())(equipmentId)
			if err != nil {
				d.Logger().WithError(err).Errorf("Unable to get equipment.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			res, err := model.Map(Transform)(model.FixedProvider(e))()
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[RestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}

func handleGetEquipmentSlots(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseEquipmentId(d.Logger(), func(equipmentId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			e, err := GetById(d.Context())(equipmentId)
			if err != nil {
				d.Logger().WithError(err).Errorf("Unable to get equipment.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			res, err := model.Map(TransformSlot)(model.FixedProvider(e))()
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[[]SlotRestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}
