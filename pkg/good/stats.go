package good

import "github.com/shizukayuki/excel-hk4e"

var (
	FightPropToGOOD map[excel.FightProp]StatKey
	GOODToFightProp = map[StatKey]excel.FightProp{
		HP:        excel.FIGHT_PROP_HP,
		HPP:       excel.FIGHT_PROP_HP_PERCENT,
		ATK:       excel.FIGHT_PROP_ATTACK,
		ATKP:      excel.FIGHT_PROP_ATTACK_PERCENT,
		DEF:       excel.FIGHT_PROP_DEFENSE,
		DEFP:      excel.FIGHT_PROP_DEFENSE_PERCENT,
		CR:        excel.FIGHT_PROP_CRITICAL,
		CD:        excel.FIGHT_PROP_CRITICAL_HURT,
		EM:        excel.FIGHT_PROP_ELEMENT_MASTERY,
		ER:        excel.FIGHT_PROP_CHARGE_EFFICIENCY,
		Heal:      excel.FIGHT_PROP_HEAL_ADD,
		PhysicalP: excel.FIGHT_PROP_PHYSICAL_ADD_HURT,
		PyroP:     excel.FIGHT_PROP_FIRE_ADD_HURT,
		HydroP:    excel.FIGHT_PROP_WATER_ADD_HURT,
		DendroP:   excel.FIGHT_PROP_GRASS_ADD_HURT,
		ElectroP:  excel.FIGHT_PROP_ELEC_ADD_HURT,
		AnemoP:    excel.FIGHT_PROP_WIND_ADD_HURT,
		CryoP:     excel.FIGHT_PROP_ICE_ADD_HURT,
		GeoP:      excel.FIGHT_PROP_ROCK_ADD_HURT,
	}
	GOODToString = map[StatKey]string{
		HP:        "HP",
		HPP:       "HP%",
		ATK:       "ATK",
		ATKP:      "ATK%",
		DEF:       "DEF",
		DEFP:      "DEF%",
		CR:        "CRIT Rate",
		CD:        "CRIT DMG",
		EM:        "Elemental Mastery",
		ER:        "Energy Recharge",
		Heal:      "Healing Bonus",
		PhysicalP: "Physical DMG Bonus",
		PyroP:     "Pyro DMG Bonus",
		HydroP:    "Hydro DMG Bonus",
		DendroP:   "Dendro DMG Bonus",
		ElectroP:  "Electro DMG Bonus",
		AnemoP:    "Anemo DMG Bonus",
		CryoP:     "Cryo DMG Bonus",
		GeoP:      "Geo DMG Bonus",
	}
)

func init() {
	FightPropToGOOD = make(map[excel.FightProp]StatKey, len(GOODToFightProp))
	for k, v := range GOODToFightProp {
		FightPropToGOOD[v] = k
	}
}

type Stats [EndStatType]float32

func (s *Stats) Merge(v *Stats) *Stats {
	for i := range s {
		s[i] += v[i]
	}
	return s
}

func (s *Stats) Get(prop StatKey) float32 {
	return s[prop]
}

func (s *Stats) Set(prop StatKey, value float32) {
	if prop > UnknownStatKey && prop < EndStatType {
		s[prop] = value
	}
}

func (s *Stats) Add(prop StatKey, value float32) {
	s.Set(prop, s.Get(prop)+value)
}

func (s *Stats) GetFightProp(prop excel.FightProp) float32 {
	return s.Get(FightPropToGOOD[prop])
}

func (s *Stats) SetFightProp(prop excel.FightProp, value float32) {
	s.Set(FightPropToGOOD[prop], value)
}

func (s *Stats) AddFightProp(prop excel.FightProp, value float32) {
	s.SetFightProp(prop, s.GetFightProp(prop)+value)
}
