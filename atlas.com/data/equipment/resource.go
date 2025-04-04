package equipment

import (
	"atlas-data/rest"
	"github.com/Chronicle20/atlas-rest/server"
	"github.com/gorilla/mux"
	"github.com/jtumidanski/api2go/jsonapi"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func InitResource(db *gorm.DB) func(si jsonapi.ServerInformation) server.RouteInitializer {
	return func(si jsonapi.ServerInformation) server.RouteInitializer {
		return func(router *mux.Router, l logrus.FieldLogger) {
			registerGet := rest.RegisterHandler(l)(si)

			r := router.PathPrefix("/data/equipment").Subrouter()
			r.HandleFunc("/{equipmentId}", registerGet("get_equipment_statistics", handleGetEquipmentStatistics(db))).Methods(http.MethodGet)
			r.HandleFunc("/{equipmentId}/slots", registerGet("get_equipment_slots", handleGetEquipmentSlots(db))).Methods(http.MethodGet)
		}
	}
}

func handleGetEquipmentStatistics(db *gorm.DB) func(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return func(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
		return rest.ParseEquipmentId(d.Logger(), func(equipmentId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				s := NewStorage(d.Logger(), db)
				res, err := s.GetById(d.Context())(strconv.Itoa(int(equipmentId)))
				if err != nil {
					d.Logger().WithError(err).Errorf("Unable to get equipment.")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				query := r.URL.Query()
				queryParams := jsonapi.ParseQueryFields(&query)
				server.MarshalResponse[RestModel](d.Logger())(w)(c.ServerInformation())(queryParams)(res)
			}
		})
	}
}

func handleGetEquipmentSlots(db *gorm.DB) func(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return func(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
		return rest.ParseEquipmentId(d.Logger(), func(equipmentId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				s := NewStorage(d.Logger(), db)
				res, err := s.GetById(d.Context())(strconv.Itoa(int(equipmentId)))
				if err != nil {
					d.Logger().WithError(err).Errorf("Unable to get equipment.")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				query := r.URL.Query()
				queryParams := jsonapi.ParseQueryFields(&query)
				server.MarshalResponse[[]SlotRestModel](d.Logger())(w)(c.ServerInformation())(queryParams)(res.EquipSlots)
			}
		})
	}
}
