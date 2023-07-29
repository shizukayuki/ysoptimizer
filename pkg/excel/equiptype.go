package excel

//go:generate enumer --json --type=EquipType $GOFILE
type EquipType uint32

const (
	EQUIP_NONE EquipType = iota
	EQUIP_BRACER
	EQUIP_NECKLACE
	EQUIP_SHOES
	EQUIP_RING
	EQUIP_DRESS
	EQUIP_WEAPON
)
