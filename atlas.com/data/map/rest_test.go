package _map

import (
	npc2 "atlas-data/npc"
	"context"
	"github.com/Chronicle20/atlas-rest/server"
	tenant "github.com/Chronicle20/atlas-tenant"
	"github.com/jtumidanski/api2go/jsonapi"
	"github.com/sirupsen/logrus/hooks/test"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

type Server struct {
	baseUrl string
	prefix  string
}

func (s Server) GetBaseURL() string {
	return s.baseUrl
}

func (s Server) GetPrefix() string {
	return s.prefix
}

func GetServer() Server {
	return Server{
		baseUrl: "",
		prefix:  "/api/",
	}
}

func TestRest(t *testing.T) {
	tt := testTenant()
	l, _ := test.NewNullLogger()
	ctx := tenant.WithContext(context.Background(), tt)

	GetMapStringRegistry().Add(tt, MapString{
		id:         strconv.Itoa(50000),
		mapName:    "Dangerous Forest",
		streetName: "Maple Road",
	})
	npc2.GetNpcStringRegistry().Add(tt, npc2.NewNpcString(strconv.Itoa(2003), "Robin"))
	npc2.GetNpcStringRegistry().Add(tt, npc2.NewNpcString(strconv.Itoa(2005), "Sam"))

	input, err := Read(l)(ctx)("", 0, fixedNodeProvider)()
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	server.MarshalResponse[RestModel](l)(rr)(GetServer())(map[string][]string{})(input)

	if rr.Code != http.StatusOK {
		t.Fatalf("Failed to write rest model: %v", err)
	}

	body := rr.Body.Bytes()

	var output RestModel
	err = jsonapi.Unmarshal(body, &output)

	ok := compare(input, output)
	if !ok {
		t.Fatalf("Failed to compare model: %v", input.Id)
	}
}

func TestLinkedRest(t *testing.T) {
	tt := testTenant()
	l, _ := test.NewNullLogger()
	ctx := tenant.WithContext(context.Background(), tt)

	GetMapStringRegistry().Add(tt, MapString{
		id:         strconv.Itoa(100020100),
		mapName:    "Henesys Pig Farm",
		streetName: "Mini Dungeon",
	})

	input, err := Read(l)(ctx)("", 0, linkedNodeProvider)()
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	server.MarshalResponse[RestModel](l)(rr)(GetServer())(map[string][]string{})(input)

	if rr.Code != http.StatusOK {
		t.Fatalf("Failed to write rest model: %v", err)
	}

	body := rr.Body.Bytes()

	var output RestModel
	err = jsonapi.Unmarshal(body, &output)

	ok := compare(input, output)
	if !ok {
		t.Fatalf("Failed to compare model: %v", input.Id)
	}
}

func compare(m1 RestModel, m2 RestModel) bool {
	return reflect.DeepEqual(m1, m2)
}
