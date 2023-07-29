package excel

var (
	WeaponCodexExcelConfigData   []*WeaponCodex
	WeaponCurveExcelConfigData   []*CurveData
	WeaponExcelConfigData        []*Weapon
	WeaponPromoteExcelConfigData []*WeaponPromote
)

func init() {
	load("ExcelBinOutput/WeaponCodexExcelConfigData.json", &WeaponCodexExcelConfigData)
	load("ExcelBinOutput/WeaponCurveExcelConfigData.json", &WeaponCurveExcelConfigData)
	load("ExcelBinOutput/WeaponExcelConfigData.json", &WeaponExcelConfigData)
	load("ExcelBinOutput/WeaponPromoteExcelConfigData.json", &WeaponPromoteExcelConfigData)
}

type WeaponCodex struct {
	WeaponId uint32
	IsDisuse bool
}

func (w *WeaponCodex) Weapon() *Weapon {
	return FindWeapon(w.WeaponId)
}

type Weapon struct {
	WeaponType string
	RankLevel  uint32
	SkillAffix []uint32
	WeaponProp []struct {
		PropType  FightProp
		InitValue float32
		Type      string
	}
	WeaponPromoteId uint32
	Id              uint32
	NameTextMapHash TextMapHash
}

func (w *Weapon) Name() string {
	return w.NameTextMapHash.String()
}

func (w *Weapon) Codex() *WeaponCodex {
	for _, v := range WeaponCodexExcelConfigData {
		if v.WeaponId == w.Id {
			return v
		}
	}
	return nil
}

func (w *Weapon) Affix(level uint32) *EquipAffix {
	if len(w.SkillAffix) == 0 {
		return nil
	}
	return FindEquipAffix(w.SkillAffix[0], level)
}

func (w *Weapon) Curve(level uint32) *CurveData {
	return FindCurveData(WeaponCurveExcelConfigData, level)
}

func (w *Weapon) Promote(level uint32) *WeaponPromote {
	for _, v := range WeaponPromoteExcelConfigData {
		if v.WeaponPromoteId == w.WeaponPromoteId && v.PromoteLevel == level {
			return v
		}
	}
	return nil
}

type WeaponPromote struct {
	WeaponPromoteId uint32
	PromoteLevel    uint32
	AddProps        []FightPropData
	UnlockMaxLevel  uint32
}

func FindWeapon(id uint32) *Weapon {
	for _, v := range WeaponExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}
