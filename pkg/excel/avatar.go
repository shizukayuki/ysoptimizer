package excel

var (
	AvatarCodexExcelConfigData      []*AvatarCodex
	AvatarCurveExcelConfigData      []*CurveData
	AvatarExcelConfigData           []*Avatar
	AvatarPromoteExcelConfigData    []*AvatarPromote
	AvatarSkillDepotExcelConfigData []*AvatarSkillDepot
)

func init() {
	load("ExcelBinOutput/AvatarCodexExcelConfigData.json", &AvatarCodexExcelConfigData)
	load("ExcelBinOutput/AvatarCurveExcelConfigData.json", &AvatarCurveExcelConfigData)
	load("ExcelBinOutput/AvatarExcelConfigData.json", &AvatarExcelConfigData)
	load("ExcelBinOutput/AvatarPromoteExcelConfigData.json", &AvatarPromoteExcelConfigData)
	load("ExcelBinOutput/AvatarSkillDepotExcelConfigData.json", &AvatarSkillDepotExcelConfigData)
}

type AvatarCodex struct {
	AvatarId uint32
}

func (a *AvatarCodex) Avatar() *Avatar {
	return FindAvatar(a.AvatarId)
}

type Avatar struct {
	UseType             string
	BodyType            string
	QualityType         string
	ChargeEfficiency    float32
	WeaponType          string
	SkillDepotId        uint32
	StaminaRecoverSpeed float32
	CandSkillDepotIds   []uint32
	AvatarPromoteId     uint32
	HpBase              float32
	AttackBase          float32
	DefenseBase         float32
	Critical            float32
	CriticalHurt        float32
	PropGrowCurves      []PropGrowCurves
	Id                  uint32
	NameTextMapHash     TextMapHash
}

func (a *Avatar) Name() string {
	return a.NameTextMapHash.String()
}

func (a *Avatar) Codex() *AvatarCodex {
	for _, v := range AvatarCodexExcelConfigData {
		if v.AvatarId == a.Id {
			return v
		}
	}
	return nil
}

func (a *Avatar) SkillDepot() *AvatarSkillDepot {
	return FindSkillDepot(a.SkillDepotId)
}

func (a *Avatar) Curve(level uint32) *CurveData {
	return FindCurveData(AvatarCurveExcelConfigData, level)
}

func (a *Avatar) Promote(level uint32) *AvatarPromote {
	for _, v := range AvatarPromoteExcelConfigData {
		if v.AvatarPromoteId == a.AvatarPromoteId && v.PromoteLevel == level {
			return v
		}
	}
	return nil
}

type AvatarPromote struct {
	AvatarPromoteId uint32
	PromoteLevel    uint32
	UnlockMaxLevel  uint32
	AddProps        []FightPropData
}

type AvatarSkillDepot struct {
	Id                      uint32
	EnergySkill             uint32
	Skills                  []uint32
	SubSkills               []uint32
	AttackModeSkill         uint32
	Talents                 []uint32
	InherentProudSkillOpens []struct {
		ProudSkillGroupId      uint32
		NeedAvatarPromoteLevel uint32
	}
}

func FindAvatar(id uint32) *Avatar {
	for _, v := range AvatarExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func FindSkillDepot(id uint32) *AvatarSkillDepot {
	for _, v := range AvatarSkillDepotExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}
