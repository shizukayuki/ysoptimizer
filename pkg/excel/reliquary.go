package excel

import "sort"

var (
	ReliquaryAffixExcelConfigData    []*ReliquaryAffix
	ReliquaryCodexExcelConfigData    []*ReliquaryCodex
	ReliquaryExcelConfigData         []*Reliquary
	ReliquaryLevelExcelConfigData    []*ReliquaryLevel
	ReliquaryMainPropExcelConfigData []*ReliquaryMainProp
	ReliquarySetExcelConfigData      []*ReliquarySet
)

func init() {
	load("ExcelBinOutput/ReliquaryAffixExcelConfigData.json", &ReliquaryAffixExcelConfigData)
	load("ExcelBinOutput/ReliquaryCodexExcelConfigData.json", &ReliquaryCodexExcelConfigData)
	load("ExcelBinOutput/ReliquaryExcelConfigData.json", &ReliquaryExcelConfigData)
	load("ExcelBinOutput/ReliquaryLevelExcelConfigData.json", &ReliquaryLevelExcelConfigData)
	load("ExcelBinOutput/ReliquaryMainPropExcelConfigData.json", &ReliquaryMainPropExcelConfigData)
	load("ExcelBinOutput/ReliquarySetExcelConfigData.json", &ReliquarySetExcelConfigData)
}

type ReliquaryAffix struct {
	Id        uint32
	DepotId   uint32
	GroupId   uint32
	PropType  FightProp
	PropValue float32
}

func (r *ReliquaryAffix) Min() *ReliquaryAffix {
	return r.Rolls()[0]
}

func (r *ReliquaryAffix) Max() *ReliquaryAffix {
	v := r.Rolls()
	return v[len(v)-1]
}

func (r *ReliquaryAffix) Rolls() []*ReliquaryAffix {
	var rolls []*ReliquaryAffix
	for _, v := range ReliquaryAffixExcelConfigData {
		if v.DepotId == r.DepotId && v.PropType == r.PropType {
			rolls = append(rolls, v)
		}
	}
	sort.Slice(rolls, func(i, j int) bool { return rolls[i].PropValue < rolls[j].PropValue })
	return rolls
}

type ReliquaryCodex struct {
	Id        uint32
	SuitId    uint32
	Level     uint32
	CupId     uint32
	LeatherId uint32
	CapId     uint32
	FlowerId  uint32
	SandId    uint32
}

func (r *ReliquaryCodex) Reliquary(typ EquipType) *Reliquary {
	return FindReliquary(r.equipType(typ))
}

func (r *ReliquaryCodex) Set() *ReliquarySet {
	return FindReliquarySet(r.SuitId)
}

func (r *ReliquaryCodex) equipType(typ EquipType) uint32 {
	switch typ {
	case EQUIP_BRACER:
		return r.FlowerId
	case EQUIP_NECKLACE:
		return r.LeatherId
	case EQUIP_SHOES:
		return r.SandId
	case EQUIP_RING:
		return r.CupId
	case EQUIP_DRESS:
		return r.CapId
	default:
		return 0
	}
}

type Reliquary struct {
	EquipType         EquipType
	RankLevel         uint32
	MainPropDepotId   uint32
	AppendPropDepotId uint32
	AppendPropNum     uint32
	SetId             uint32
	MaxLevel          uint32
	Id                uint32
	NameTextMapHash   TextMapHash
}

func (r *Reliquary) Name() string {
	return r.NameTextMapHash.String()
}

func (r *Reliquary) Codex() *ReliquaryCodex {
	for _, v := range ReliquaryCodexExcelConfigData {
		if v.equipType(r.EquipType) == r.Id {
			return v
		}
	}
	return nil
}

func (r *Reliquary) Set() *ReliquarySet {
	return FindReliquarySet(r.SetId)
}

func (r *Reliquary) Level(level uint32) *ReliquaryLevel {
	for _, v := range ReliquaryLevelExcelConfigData {
		if v.Rank == r.RankLevel && v.Level == level {
			return v
		}
	}
	return nil
}

func (r *Reliquary) MainProp(id uint32) *ReliquaryMainProp {
	for _, v := range ReliquaryMainPropExcelConfigData {
		if v.PropDepotId == r.MainPropDepotId && v.Id == id {
			return v
		}
	}
	return nil
}

func (r *Reliquary) AppendProp(id uint32) *ReliquaryAffix {
	for _, v := range ReliquaryAffixExcelConfigData {
		if v.DepotId == r.AppendPropDepotId && v.Id == id {
			return v
		}
	}
	return nil
}

type ReliquaryLevel struct {
	Rank     uint32
	Level    uint32
	AddProps []FightPropData
}

func (r *ReliquaryLevel) Stat(prop FightProp) float32 {
	for _, v := range r.AddProps {
		if v.PropType == prop {
			return v.Value
		}
	}
	return 0
}

type ReliquaryMainProp struct {
	Id          uint32
	PropDepotId uint32
	PropType    FightProp
}

type ReliquarySet struct {
	SetId        uint32
	SetNeedNum   []uint32
	EquipAffixId uint32
}

func (r *ReliquarySet) Codex(level uint32) *ReliquaryCodex {
	for _, v := range ReliquaryCodexExcelConfigData {
		if v.SuitId == r.SetId && v.Level == level {
			return v
		}
	}
	return nil
}

func (r *ReliquarySet) Affix(level uint32) *EquipAffix {
	return FindEquipAffix(r.EquipAffixId, level)
}

func FindReliquary(id uint32) *Reliquary {
	for _, v := range ReliquaryExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func FindReliquarySet(id uint32) *ReliquarySet {
	for _, v := range ReliquarySetExcelConfigData {
		if v.SetId == id {
			return v
		}
	}
	return nil
}
