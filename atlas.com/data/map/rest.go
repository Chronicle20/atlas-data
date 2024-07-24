package _map

import (
	"github.com/Chronicle20/atlas-model/model"
)

type RestModel struct {
	Id                string                    `json:"-"`
	Name              string                    `json:"name"`
	StreetName        string                    `json:"streetName"`
	ReturnMapId       uint32                    `json:"returnMapId"`
	MonsterRate       float64                   `json:"monsterRate"`
	OnFirstUserEnter  string                    `json:"onFirstUserEnter"`
	OnUserEnter       string                    `json:"onUserEnter"`
	FieldLimit        uint32                    `json:"fieldLimit"`
	MobInterval       uint32                    `json:"mobInterval"`
	Seats             uint32                    `json:"seats"`
	Clock             bool                      `json:"clock"`
	EverLast          bool                      `json:"everLast"`
	Town              bool                      `json:"town"`
	DecHP             uint32                    `json:"decHP"`
	ProtectItem       uint32                    `json:"protectItem"`
	ForcedReturnMapId uint32                    `json:"forcedReturnMapId"`
	Boat              bool                      `json:"boat"`
	TimeLimit         int32                     `json:"timeLimit"`
	FieldType         uint32                    `json:"fieldType"`
	MobCapacity       uint32                    `json:"mobCapacity"`
	Recovery          float64                   `json:"recovery"`
	MapArea           RectangleRestModel        `json:"mapArea"`
	Areas             []RectangleRestModel      `json:"areas"`
	BackgroundTypes   []BackgroundTypeRestModel `json:"backgroundTypes"`
}

func (r RestModel) GetName() string {
	return "maps"
}

func (r RestModel) GetID() string {
	return r.Id
}

func Transform(m Model) (RestModel, error) {
	ma, err := TransformRectangle(m.mapArea)
	if err != nil {
		return RestModel{}, err
	}

	as, err := model.SliceMap(model.FixedProvider(m.areas), TransformRectangle)()
	if err != nil {
		return RestModel{}, err
	}

	bt, err := model.SliceMap(model.FixedProvider(m.backgroundTypes), TransformBackgroundType)()
	if err != nil {
		return RestModel{}, err
	}

	return RestModel{
		Name:              m.name,
		StreetName:        m.streetName,
		ReturnMapId:       m.returnMapId,
		MonsterRate:       m.monsterRate,
		OnFirstUserEnter:  m.onFirstUserEnter,
		OnUserEnter:       m.onUserEnter,
		FieldLimit:        m.fieldLimit,
		MobInterval:       m.mobInterval,
		Seats:             m.seats,
		Clock:             m.clock,
		EverLast:          m.everLast,
		Town:              m.town,
		DecHP:             m.decHp,
		ProtectItem:       m.protectItem,
		ForcedReturnMapId: m.forcedReturnMapId,
		Boat:              m.boat,
		TimeLimit:         m.timeLimit,
		FieldType:         m.fieldType,
		MobCapacity:       m.mobCapacity,
		Recovery:          m.recovery,
		MapArea:           ma,
		Areas:             as,
		BackgroundTypes:   bt,
	}, nil
}

type RectangleRestModel struct {
	X      int16 `json:"x"`
	Y      int16 `json:"y"`
	Width  int16 `json:"width"`
	Height int16 `json:"height"`
}

func TransformRectangle(m Rectangle) (RectangleRestModel, error) {
	return RectangleRestModel{
		X:      m.x,
		Y:      m.y,
		Width:  m.width,
		Height: m.height,
	}, nil
}

type BackgroundTypeRestModel struct {
	LayerNumber    uint32 `json:"layerNumber"`
	BackgroundType uint32 `json:"backgroundType"`
}

func TransformBackgroundType(m BackgroundType) (BackgroundTypeRestModel, error) {
	return BackgroundTypeRestModel{
		LayerNumber:    m.layerNumber,
		BackgroundType: m.backgroundType,
	}, nil
}
