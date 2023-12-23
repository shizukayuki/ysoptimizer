package good

import (
	"fmt"
	"math"
	"strings"

	"github.com/shizukayuki/ysoptimizer/pkg/excel"
)

type Datebase struct {
	Format     string            `json:"format"`
	Version    int               `json:"version"`
	Source     string            `json:"source"`
	Characters []*Character      `json:"characters,omitempty"`
	Artifacts  []*Artifact       `json:"artifacts,omitempty"`
	Weapons    []*Weapon         `json:"weapons,omitempty"`
	Materials  map[string]uint32 `json:"materials,omitempty"`
}

type Character struct {
	Key           CharacterKey `json:"key"`
	Level         uint32       `json:"level"`
	Ascension     uint32       `json:"ascension"`
	Constellation uint32       `json:"constellation"`
	Talent        Talent       `json:"talent"`
	Favorite      bool         `json:"favorite"`
}

type Talent struct {
	Auto  uint32 `json:"auto"`
	Skill uint32 `json:"skill"`
	Burst uint32 `json:"burst"`
}

type Artifact struct {
	Stats `json:"-"`

	SetKey      ArtifactSetKey `json:"setKey"`
	Rarity      uint32         `json:"rarity"`
	Level       uint32         `json:"level"`
	SlotKey     SlotKey        `json:"slotKey"`
	MainStatKey StatKey        `json:"mainStatKey"`
	Substats    []Affix        `json:"substats"`
	Location    CharacterKey   `json:"location"`
	Lock        bool           `json:"lock"`
}

func (a Artifact) String() string {
	mainAffix := GOODToString[a.MainStatKey]
	mainValue := a.Get(a.MainStatKey)
	if a.MainStatKey.PercentStat() {
		mainAffix += fmt.Sprintf("+%.1f%%", mainValue*100)
	} else {
		mainAffix += fmt.Sprintf("+%.0f", mainValue)
	}

	var affixes string
	for _, s := range a.Substats {
		var buf strings.Builder

		buf.WriteString(GOODToString[s.Key])
		if s.Key.PercentStat() {
			fmt.Fprintf(&buf, "+%.1f%%", s.Value)
		} else {
			fmt.Fprintf(&buf, "+%.0f", s.Value)
		}

		affixes += fmt.Sprintf("[%s]%s", buf.String(), strings.Repeat(" ", 28-buf.Len()))
	}
	affixes = strings.TrimSpace(affixes)

	var sb strings.Builder

	// fmt.Fprintf(&sb, "[%s] ", strings.ToUpper(string(a.SlotKey.String()[0])))
	fmt.Fprintf(&sb, "%-36s", a.SetKey)
	fmt.Fprintf(&sb, "%-8s", fmt.Sprintf("%4.1f%%", a.CritValue()*100))
	fmt.Fprintf(&sb, "%-30s", mainAffix)
	sb.WriteString(affixes)

	return sb.String()
}

func (a *Artifact) CritValue() float32 {
	switch a.MainStatKey {
	case CR:
		return a.Get(CD)
	case CD:
		return a.Get(CR) * 2
	default:
		return a.Get(CR)*2 + a.Get(CD)
	}
}

func (a *Artifact) Process() error {
	a.Stats = Stats{}
	relic := excel.FindReliquarySet(uint32(a.SetKey)).Codex(a.Rarity).Reliquary(a.SlotKey.EquipType())
	mainStat := relic.Level(a.Level + 1).Stat(GOODToFightProp[a.MainStatKey])
	a.Stats.Set(a.MainStatKey, mainStat)
	for _, s := range a.Substats {
		v := s.Value
		if s.Key.PercentStat() {
			v /= 100
		}
		v = correctAffix(affixPermutations(relic, GOODToFightProp[s.Key]), v)
		a.Stats.Set(s.Key, v)
	}
	return nil
}

func affixPermutations(r *excel.Reliquary, prop excel.FightProp) []float32 {
	var perm []float32
	rolls := affixRolls(r, prop)
	for i := 1; i <= (int(r.MaxLevel)/4)+1; i++ {
		perm = append(perm, choose(0, rolls, i)...)
	}
	return perm
}

func affixRolls(r *excel.Reliquary, prop excel.FightProp) []float32 {
	var rolls []float32
	for _, d := range excel.ReliquaryAffixExcelConfigData {
		if d.DepotId == r.AppendPropDepotId && d.PropType == prop {
			rolls = append(rolls, d.PropValue)
		}
	}
	return rolls
}

func choose(prefix float32, arr []float32, k int) []float32 {
	if k == 0 {
		return []float32{prefix}
	}

	var result []float32
	for i, v := range arr {
		result = append(result, choose(prefix+v, arr[i:], k-1)...)
	}
	return result
}

func correctAffix(perm []float32, propValue float32) float32 {
	var result float32

	diff := math.MaxFloat32
	for _, v := range perm {
		rem := math.Abs(float64(v - propValue))
		if rem < diff {
			diff = rem
			result = v
		}
	}
	return result
}

type Affix struct {
	Key   StatKey `json:"key"`
	Value float32 `json:"value"`
}

type Weapon struct {
	Key        WeaponKey    `json:"key"`
	Level      uint32       `json:"level"`
	Ascension  uint32       `json:"ascension"`
	Refinement uint32       `json:"refinement"`
	Location   CharacterKey `json:"location"`
	Lock       bool         `json:"lock"`
}
