package good

import (
	"fmt"

	"github.com/shizukayuki/excel-hk4e"
)

//go:generate enumer --json --linecomment --type=SlotKey $GOFILE
type SlotKey int

const (
	Flower  SlotKey = iota // flower
	Plume                  // plume
	Sands                  // sands
	Goblet                 // goblet
	Circlet                // circlet
)

func (k SlotKey) EquipType() excel.EquipType {
	switch k {
	case Flower:
		return excel.EQUIP_BRACER
	case Plume:
		return excel.EQUIP_NECKLACE
	case Sands:
		return excel.EQUIP_SHOES
	case Goblet:
		return excel.EQUIP_RING
	case Circlet:
		return excel.EQUIP_DRESS
	}
	panic(fmt.Errorf("unknown slot key: %s", k))
}

//go:generate enumer --json --linecomment --type=StatKey $GOFILE
type StatKey int

const (
	UnknownStatKey StatKey = iota //
	HP                            // hp
	HPP                           // hp_
	ATK                           // atk
	ATKP                          // atk_
	DEF                           // def
	DEFP                          // def_
	CR                            // critRate_
	CD                            // critDMG_
	EM                            // eleMas
	ER                            // enerRech_
	Heal                          // heal_
	PhysicalP                     // physical_dmg_
	PyroP                         // pyro_dmg_
	HydroP                        // hydro_dmg_
	DendroP                       // dendro_dmg_
	ElectroP                      // electro_dmg_
	AnemoP                        // anemo_dmg_
	CryoP                         // cryo_dmg_
	GeoP                          // geo_dmg_
	EndStatType
)

func (k StatKey) PercentStat() bool {
	s := k.String()
	return s[len(s)-1:] == "_"
}

//go:generate go run ./gen.go $GOPACKAGE keys_gen.go
//go:generate enumer --json --linecomment --type=ArtifactSetKey keys_gen.go
//go:generate enumer --json --linecomment --type=CharacterKey keys_gen.go
//go:generate enumer --json --linecomment --type=WeaponKey keys_gen.go
