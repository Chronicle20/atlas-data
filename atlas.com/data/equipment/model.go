package equipment

type Model struct {
	itemId        uint32
	strength      uint16
	dexterity     uint16
	intelligence  uint16
	luck          uint16
	weaponAttack  uint16
	weaponDefense uint16
	magicAttack   uint16
	magicDefense  uint16
	accuracy      uint16
	avoidability  uint16
	speed         uint16
	jump          uint16
	hp            uint16
	mp            uint16
	slots         uint16
	cash          bool
	slotName      string
	slotWz        string
	slotIndex     []int16
}

func (m Model) Id() uint32 {
	return m.itemId
}
