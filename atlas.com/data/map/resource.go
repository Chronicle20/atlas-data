package _map

import (
	point2 "atlas-data/map/point"
	"atlas-data/monster"
	"atlas-data/npc"
	"atlas-data/point"
	"atlas-data/portal"
	"atlas-data/reactor"
	"atlas-data/rest"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/Chronicle20/atlas-rest/server"
	"github.com/gorilla/mux"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const (
	getMap               = "get_map"
	getMapPortalsByName  = "get_map_portals_by_name"
	getMapPortals        = "get_map_portals"
	getMapPortal         = "get_map_portal"
	getMapReactors       = "get_map_reactors"
	getMapNPCsByObjectId = "get_map_npcs_by_object_id"
	getMapNPCs           = "get_map_npcs"
	getMapNPC            = "get_map_npc"
	getMapMonsters       = "get_map_monsters"
	getMapDropPosition   = "get_map_drop_position"
)

func InitResource(si jsonapi.ServerInformation) server.RouteInitializer {
	return func(router *mux.Router, l logrus.FieldLogger) {
		registerGet := rest.RegisterHandler(l)(si)
		registerInput := rest.RegisterInputHandler[DropPositionRestModel](l)(si)

		r := router.PathPrefix("/maps").Subrouter()
		r.HandleFunc("/{mapId}", registerGet(getMap, handleGetMapRequest)).Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/portals", registerGet(getMapPortalsByName, handleGetMapPortalsByNameRequest)).Queries("name", "{name}").Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/portals", registerGet(getMapPortals, handleGetMapPortalsRequest)).Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/portals/{portalId}", registerGet(getMapPortal, handleGetMapPortalRequest)).Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/reactors", registerGet(getMapReactors, handleGetMapReactorsRequest)).Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/npcs", registerGet(getMapNPCsByObjectId, handleGetMapNPCsByObjectIdRequest)).Queries("objectId", "{objectId}").Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/npcs", registerGet(getMapNPCs, handleGetMapNPCsRequest)).Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/npcs/{npcId}", registerGet(getMapNPC, handleGetMapNPCRequest)).Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/monsters", registerGet(getMapMonsters, handleGetMapMonstersRequest)).Methods(http.MethodGet)
		r.HandleFunc("/{mapId}/dropPosition", registerInput(getMapDropPosition, handleGetMapDropPositionRequest)).Methods(http.MethodPost)
	}
}

func handleGetMapRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			m, err := GetById(d.Logger(), d.Span(), c.Tenant())(mapId)
			if err != nil {
				d.Logger().WithError(err).Debugf("Unable to locate map %d.", mapId)
				w.WriteHeader(http.StatusNotFound)
				return
			}

			res, err := model.Transform(m, Transform)
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[RestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}

func handleGetMapPortalsRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ps, err := GetPortals(d.Logger(), d.Span(), c.Tenant())(mapId)
			if err != nil {
				d.Logger().WithError(err).Debugf("Unable to locate map %d.", mapId)
				w.WriteHeader(http.StatusNotFound)
				return
			}

			res, err := model.TransformAll(ps, portal.Transform)
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[[]portal.RestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}

func handleGetMapPortalsByNameRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			portalName := vars["name"]

			ps, err := GetPortalsByName(d.Logger(), d.Span(), c.Tenant())(mapId, portalName)
			if err != nil {
				d.Logger().WithError(err).Debugf("Unable to locate map %d.", mapId)
				w.WriteHeader(http.StatusNotFound)
				return
			}

			res, err := model.TransformAll(ps, portal.Transform)
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[[]portal.RestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}

func handleGetMapPortalRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return rest.ParsePortalId(d.Logger(), func(portalId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				p, err := GetPortalById(d.Logger(), d.Span(), c.Tenant())(mapId, portalId)
				if err != nil {
					d.Logger().WithError(err).Debugf("Unable to locate map %d.", mapId)
					w.WriteHeader(http.StatusNotFound)
					return
				}

				res, err := model.Transform(p, portal.Transform)
				if err != nil {
					d.Logger().WithError(err).Errorf("Creating REST model.")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				server.Marshal[portal.RestModel](d.Logger())(w)(c.ServerInformation())(res)
			}
		})
	})
}

func handleGetMapReactorsRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			rs, err := GetReactors(d.Logger(), d.Span(), c.Tenant())(mapId)
			if err != nil {
				d.Logger().WithError(err).Debugf("Unable to locate map %d.", mapId)
				w.WriteHeader(http.StatusNotFound)
				return
			}

			res, err := model.TransformAll(rs, reactor.Transform)
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[[]reactor.RestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}

func handleGetMapNPCsByObjectIdRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			objectId, err := strconv.Atoi(vars["objectId"])
			if err != nil {
				d.Logger().WithError(err).Errorf("Error parsing objectId as uint32")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			ns, err := GetNpcsByObjectId(d.Logger(), d.Span(), c.Tenant())(mapId, uint32(objectId))
			res, err := model.TransformAll(ns, npc.Transform)
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[[]npc.RestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}

func handleGetMapNPCsRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ns, err := GetNpcs(d.Logger(), d.Span(), c.Tenant())(mapId)
			res, err := model.TransformAll(ns, npc.Transform)
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[[]npc.RestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}

func handleGetMapNPCRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return rest.ParseNPC(d.Logger(), func(npcId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				ns, err := GetNpc(d.Logger(), d.Span(), c.Tenant())(mapId, npcId)
				res, err := model.Transform(ns, npc.Transform)
				if err != nil {
					d.Logger().WithError(err).Errorf("Creating REST model.")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				server.Marshal[npc.RestModel](d.Logger())(w)(c.ServerInformation())(res)
			}
		})
	})
}

func handleGetMapMonstersRequest(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ms, err := GetMonsters(d.Logger(), d.Span(), c.Tenant())(mapId)
			res, err := model.TransformAll(ms, monster.Transform)
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[[]monster.RestModel](d.Logger())(w)(c.ServerInformation())(res)
		}
	})
}

func handleGetMapDropPositionRequest(d *rest.HandlerDependency, c *rest.HandlerContext, input DropPositionRestModel) http.HandlerFunc {
	return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			p := calcDropPos(c.Tenant(), mapId, point.NewModel(input.InitialX, input.InitialY), point.NewModel(input.FallbackX, input.FallbackY))
			res, err := model.Transform(*p, point2.Transform)
			if err != nil {
				d.Logger().WithError(err).Errorf("Creating REST model.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			server.Marshal[point2.RestModel](d.Logger())(w)(c.ServerInformation())(res)
			w.WriteHeader(http.StatusCreated)
		}
	})
}

type DropPositionRestModel struct {
	InitialX  int16 `json:"initialX"`
	InitialY  int16 `json:"initialY"`
	FallbackX int16 `json:"fallbackX"`
	FallbackY int16 `json:"fallbackY"`
}
