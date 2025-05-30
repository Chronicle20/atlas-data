package equipment

import (
	"atlas-data/xml"
	"fmt"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strconv"
	"strings"
)

func parseItemId(filePath string) (uint32, error) {
	baseName := filepath.Base(filePath)
	if !strings.HasSuffix(baseName, ".img") {
		return 0, fmt.Errorf("file does not match expected format: %s", filePath)
	}
	idStr := strings.TrimSuffix(baseName, ".img")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return uint32(id), nil

}

func Read(l logrus.FieldLogger) func(np model.Provider[xml.Node]) model.Provider[RestModel] {
	return func(np model.Provider[xml.Node]) model.Provider[RestModel] {
		exml, err := np()
		itemId, err := parseItemId(exml.Name)
		if err != nil {
			return model.ErrorProvider[RestModel](err)
		}

		info, err := exml.ChildByName("info")
		if err != nil {
			info, err = exml.ChildByName("0" + strconv.Itoa(int(itemId)))
			if err != nil {
				return model.ErrorProvider[RestModel](err)
			}
			info, err = info.ChildByName("info")
			if err != nil {
				return model.ErrorProvider[RestModel](err)
			}
		}
		if info == nil {
			return model.FixedProvider(RestModel{Id: itemId})
		}

		slotStr := info.GetString("islot", "")
		slotName := getNameFromWz(slotStr)
		slotIndexes := getSlotsFromWz(slotStr)

		srm := make([]SlotRestModel, 0)
		for _, idx := range slotIndexes {
			srm = append(srm, SlotRestModel{
				Id:   slotName,
				Name: slotName,
				WZ:   slotStr,
				Slot: idx,
			})
		}

		m := RestModel{
			Id:            itemId,
			Strength:      info.GetShort("incSTR", 0),
			Dexterity:     info.GetShort("incDEX", 0),
			Intelligence:  info.GetShort("incINT", 0),
			Luck:          info.GetShort("incLUK", 0),
			WeaponAttack:  info.GetShort("incPAD", 0),
			WeaponDefense: info.GetShort("incPDD", 0),
			MagicAttack:   info.GetShort("incMAD", 0),
			MagicDefense:  info.GetShort("incMDD", 0),
			Accuracy:      info.GetShort("incACC", 0),
			Avoidability:  info.GetShort("incEVA", 0),
			Speed:         info.GetShort("incSpeed", 0),
			Jump:          info.GetShort("incJump", 0),
			HP:            info.GetShort("incMHP", 0),
			MP:            info.GetShort("incMMP", 0),
			Slots:         info.GetShort("tuc", 0),
			Cash:          info.GetBool("cash", false),
			Price:         uint32(info.GetIntegerWithDefault("price", 0)),
			EquipSlots:    srm,
		}
		return model.FixedProvider(m)
	}
}

func getSlotsFromWz(wz string) []int16 {
	switch wz {
	case "Cp":
		return []int16{-1}
	case "HrCp":
		return []int16{-1}
	case "Af":
		return []int16{-2}
	case "Ay":
		return []int16{-3}
	case "Ae":
		return []int16{-4}
	case "Ma":
		return []int16{-5}
	case "MaPn":
		return []int16{-5}
	case "Pn":
		return []int16{-6}
	case "So":
		return []int16{-7}
	case "GlGw":
		return []int16{-8}
	case "Gv":
		return []int16{-8}
	case "Sr":
		return []int16{-9}
	case "Si":
		return []int16{-10}
	case "Wp":
		return []int16{-11}
	case "WpSi":
		return []int16{-11}
	case "WpSp":
		return []int16{-11}
	case "Ri":
		return []int16{-12, -13, -15, -16}
	case "Pe":
		return []int16{-17}
	case "Tm":
		return []int16{-18}
	case "Sd":
		return []int16{-19}
	case "Me":
		return []int16{-49}
	case "Be":
		return []int16{-50}
	default:
		return []int16{0}
	}
}

func getNameFromWz(wz string) string {
	switch wz {
	case "Cp":
		return "HAT"
	case "HrCp":
		return "SPECIAL_HAT"
	case "Af":
		return "FACE_ACCESSORY"
	case "Ay":
		return "EYE_ACCESSORY"
	case "Ae":
		return "EARRINGS"
	case "Ma":
		return "TOP"
	case "MaPn":
		return "OVERALL"
	case "Pn":
		return "PANTS"
	case "So":
		return "SHOES"
	case "GlGw":
		return "GLOVES"
	case "Gv":
		return "CASH_GLOVES"
	case "Sr":
		return "CAPE"
	case "Si":
		return "SHIELD"
	case "Wp":
		return "WEAPON"
	case "WpSi":
		return "WEAPON_2"
	case "WpSp":
		return "LOW_WEAPON"
	case "Ri":
		return "RING"
	case "Pe":
		return "PENDANT"
	case "Tm":
		return "TAMED_MOB"
	case "Sd":
		return "SADDLE"
	case "Me":
		return "MEDAL"
	case "Be":
		return "BELT"
	default:
		return "PET_EQUIP"
	}
}
