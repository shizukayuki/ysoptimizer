package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/shizukayuki/excel-hk4e"
	"github.com/shizukayuki/ysoptimizer/pkg/good"
)

type ArtifactFilter func(*good.Artifact) bool

type OptimizeTarget struct {
	*good.Datebase
	Character    *good.Character
	Weapon       *good.Weapon
	CurArtifacts [5]*good.Artifact

	IgnoreEnemy bool
	Filter      ArtifactFilter
	Buffs       func(*OptimizeTarget, *OptimizeState) bool
	Target      func(*OptimizeTarget, *OptimizeState) float32
}

func (t *OptimizeTarget) calcBaseStats() OptimizeState {
	s := OptimizeState{}
	s.AddCharacterStats(t.Character)
	s.AddWeaponStats(t.Weapon)
	return s
}

func (t *OptimizeTarget) CurrentState() OptimizeState {
	state := t.calcBaseStats()
	state.Build = t.CurArtifacts
	state.calcArtifacts()
	buffs(t, &state)
	state.Result = t.Target(t, &state)
	if !t.IgnoreEnemy {
		state.Result *= enemyMult(0, 0, 0)
	}
	return state
}

func (t *OptimizeTarget) Optimize() OptimizeState {
	state := t.calcBaseStats()
	state = t.permute(state)
	if !t.IgnoreEnemy {
		state.Result *= enemyMult(0, 0, 0)
	}
	return state
}

func (t *OptimizeTarget) permute(state OptimizeState) OptimizeState {
	var wg sync.WaitGroup
	queue := make(chan OptimizeState)
	results := make([]OptimizeState, max(1, runtime.NumCPU()-1))
	slot := good.SlotKey(0)

	for i := range results {
		result := &results[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			for cpy := range queue {
				cpy = t.permuteSlot(cpy, slot+1)
				if cpy.Result > result.Result {
					*result = cpy
				}
			}
		}()
	}

	t.iter(slot, func(a *good.Artifact) {
		cpy := state
		cpy.Build[slot] = a
		cpy.Merge(&a.Stats)
		queue <- cpy
	})

	close(queue)
	wg.Wait()

	var result OptimizeState
	for _, cpy := range results {
		if cpy.Result > result.Result {
			result = cpy
		}
	}
	return result
}

func (t *OptimizeTarget) permuteSlot(state OptimizeState, slot good.SlotKey) OptimizeState {
	var result, cpy OptimizeState
	var inner func(OptimizeState, good.SlotKey)

	inner = func(state OptimizeState, slot good.SlotKey) {
		if slot >= good.SlotKey(len(state.Build)) {
			if !buffs(t, &cpy) {
				return
			}
			cpy.Result = t.Target(t, &cpy)
			if cpy.Result > result.Result {
				result = cpy
			}
			return
		}

		t.iter(slot, func(a *good.Artifact) {
			cpy = state
			cpy.Build[slot] = a
			cpy.Merge(&a.Stats)

			num := 1
			for i := good.SlotKey(0); i < slot; i++ {
				x := cpy.Build[i]
				if a.SetKey == x.SetKey {
					num++
				}
			}
			cpy.AddSetBonus(a.SetKey, num)

			inner(cpy, slot+1)
		})
	}
	inner(state, slot)

	return result
}

func (t *OptimizeTarget) iter(slot good.SlotKey, fn func(a *good.Artifact)) {
	for _, a := range t.Datebase.Artifacts {
		if a.SlotKey != slot || a.Level != 20 || a.Location != good.UnknownCharacterKey {
			continue
		}
		if t.Filter(a) {
			fn(a)
		}
	}
}

type ExtraStats struct {
	AllDMG     float32
	NormalDMG  float32
	ChargedDMG float32
	BurstCR    float32
	BurstDMG   float32
	SkillCR    float32
	SkillDMG   float32
}

type OptimizeState struct {
	Result float32
	Build  [5]*good.Artifact

	good.Stats
	ExtraStats
	BaseHP   float32
	BaseATK  float32
	BaseDEF  float32
	SetBonus good.ArtifactSetKey
}

func (s *OptimizeState) TotalATK() float32 {
	return s.BaseATK*(1+s.Stats[good.ATKP]) + s.Stats[good.ATK]
}

func (s *OptimizeState) TotalHP() float32 {
	return s.BaseHP*(1+s.Stats[good.HPP]) + s.Stats[good.HP]
}

func (s *OptimizeState) TotalDEF() float32 {
	return s.BaseDEF*(1+s.Stats[good.DEFP]) + s.Stats[good.DEF]
}

func (s *OptimizeState) CritAverage(rate, dmg float32) float32 {
	rate += s.Stats[good.CR]
	dmg += s.Stats[good.CD]
	if rate < 0 {
		rate = 0
	} else if rate > 1 {
		rate = 1
	}
	return 1 + rate*dmg
}

func (s *OptimizeState) calcArtifacts() {
	hist := map[good.ArtifactSetKey]int{}
	for _, a := range s.Build {
		if a == nil {
			continue
		}
		set := a.SetKey
		hist[set]++
		s.Merge(&a.Stats)
		s.AddSetBonus(set, hist[set])
	}
}

func (s *OptimizeState) AddCharacterStats(g *good.Character) {
	e := excel.FindAvatar(uint32(g.Key))
	s.Add(good.CR, e.Critical)
	s.Add(good.CD, e.CriticalHurt)
	s.Add(good.ER, e.ChargeEfficiency)

	curve := e.Curve(g.Level)
	for _, p := range e.PropGrowCurves {
		_, v := curve.Type(p.GrowCurve)
		switch p.Type {
		case excel.FIGHT_PROP_BASE_HP:
			s.BaseHP += e.HpBase * v
		case excel.FIGHT_PROP_BASE_ATTACK:
			s.BaseATK += e.AttackBase * v
		case excel.FIGHT_PROP_BASE_DEFENSE:
			s.BaseDEF += e.DefenseBase * v
		}
	}
	s.AddProps(e.Promote(g.Ascension).AddProps)
}

func (s *OptimizeState) AddWeaponStats(g *good.Weapon) {
	e := excel.FindWeapon(uint32(g.Key))

	curve := e.Curve(g.Level)
	for _, p := range e.WeaponProp {
		_, v := curve.Type(p.Type)
		v *= p.InitValue
		switch p.PropType {
		case excel.FIGHT_PROP_BASE_ATTACK:
			s.BaseATK += v
		default:
			s.AddFightProp(p.PropType, v)
		}
	}
	s.AddProps(e.Promote(g.Ascension).AddProps)
}

func (s *OptimizeState) AddProps(props []excel.FightPropData) {
	for _, p := range props {
		switch p.PropType {
		case excel.FIGHT_PROP_BASE_HP:
			s.BaseHP += p.Value
		case excel.FIGHT_PROP_BASE_ATTACK:
			s.BaseATK += p.Value
		case excel.FIGHT_PROP_BASE_DEFENSE:
			s.BaseDEF += p.Value
		default:
			s.AddFightProp(p.PropType, p.Value)
		}
	}
}

func (s *OptimizeState) AddSetBonus(set good.ArtifactSetKey, num int) {
	switch num {
	case 2:
		s.add2SetBonus(set)
	case 4:
		s.SetBonus = set
		s.add4SetBonus(set)
	}
}

func (s *OptimizeState) add2SetBonus(set good.ArtifactSetKey) {
	switch set {
	case good.BlizzardStrayer:
		s.Add(good.CryoP, .15)
	case good.MaidenBeloved:
		s.Add(good.Heal, .15)
	case good.GladiatorsFinale:
		s.Add(good.ATKP, .18)
	case good.ViridescentVenerer:
		s.Add(good.AnemoP, .15)
	case good.WanderersTroupe:
		s.Add(good.EM, 80)
	case good.ThunderingFury:
		s.Add(good.ElectroP, .15)
	case good.CrimsonWitchOfFlames:
		s.Add(good.PyroP, .15)
	case good.NoblesseOblige:
		s.BurstDMG += .20
	case good.BloodstainedChivalry:
		s.Add(good.PhysicalP, .25)
	case good.ArchaicPetra:
		s.Add(good.GeoP, .15)
	case good.HeartOfDepth:
		s.Add(good.HydroP, .15)
	case good.TenacityOfTheMillelith:
		s.Add(good.HPP, .20)
	case good.PaleFlame:
		s.Add(good.PhysicalP, .25)
	case good.ShimenawasReminiscence:
		s.Add(good.ATKP, .18)
	case good.EmblemOfSeveredFate:
		s.Add(good.ER, .20)
	case good.HuskOfOpulentDreams:
		s.Add(good.DEFP, .30)
	case good.OceanHuedClam:
		s.Add(good.Heal, .15)
	case good.VermillionHereafter:
		s.Add(good.ATKP, .18)
	case good.EchoesOfAnOffering:
		s.Add(good.ATKP, .18)
	case good.DeepwoodMemories:
		s.Add(good.DendroP, .15)
	case good.GildedDreams:
		s.Add(good.EM, 80)
	case good.DesertPavilionChronicle:
		s.Add(good.AnemoP, .15)
	case good.FlowerOfParadiseLost:
		s.Add(good.EM, 80)
	case good.NymphsDream:
		s.Add(good.HydroP, .15)
	case good.VourukashasGlow:
		s.Add(good.HPP, .20)
	case good.MarechausseeHunter:
		s.NormalDMG += .15
		s.ChargedDMG += .15
	case good.GoldenTroupe:
		s.SkillDMG += .20
	case good.SongOfDaysPast:
		s.Add(good.Heal, .15)
	case good.NighttimeWhispersInTheEchoingWoods:
		s.Add(good.ATKP, .18)
	case good.FragmentOfHarmonicWhimsy:
		s.Add(good.ATKP, .18)
	case good.UnfinishedReverie:
		s.Add(good.ATKP, .18)
	}
}

func (s *OptimizeState) add4SetBonus(set good.ArtifactSetKey) {
	switch set {
	case good.EmblemOfSeveredFate:
		v := s.Get(good.ER) * .25
		if v > .75 {
			v = .75
		}
		s.BurstDMG += v
	case good.GoldenTroupe:
		s.SkillDMG += .25
	}
}

func (s OptimizeState) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "[ HP: %.0f(+%.0f) / %.1f%% ]\n", s.TotalHP(), s.TotalHP()-s.BaseHP, (s.TotalHP()/s.BaseHP-1)*100)
	fmt.Fprintf(&sb, "[ ATK: %.0f(+%.0f) / %.1f%% ]\n", s.TotalATK(), s.TotalATK()-s.BaseATK, (s.TotalATK()/s.BaseATK-1)*100)
	fmt.Fprintf(&sb, "[ DEF: %.0f(+%.0f) / %.1f%% ]\n", s.TotalDEF(), s.TotalDEF()-s.BaseDEF, (s.TotalDEF()/s.BaseDEF-1)*100)

	if s.Get(good.EM) > 0 {
		fmt.Fprintf(&sb, "[ Elemental Mastery: %.0f ]\n", s.Get(good.EM))
	}

	if s.Get(good.Heal) > 0 {
		fmt.Fprintf(&sb, "[ Healing Bonus: %.1f%% ]\n", s.Get(good.Heal)*100)
	}

	fmt.Fprintf(&sb, "[ Energy Recharge: %.1f%% ]\n", s.Get(good.ER)*100)

	if v := s.Build[good.Goblet].MainStatKey; strings.HasSuffix(v.String(), "_dmg_") {
		fmt.Fprintf(&sb, "[ %s: %.1f%% ]\n", good.GOODToString[v], s.Get(v)*100)
	}

	var cv float32
	for _, a := range s.Build {
		cv += a.Get(good.CR)*2 + a.Get(good.CD)
	}
	fmt.Fprintf(&sb, "[ CRIT: %.1f%%:%.1f%% ]\n", s.Get(good.CR)*100, s.Get(good.CD)*100)
	fmt.Fprintf(&sb, "[ CV: %.1f%% / CP: %.1f%% ]\n", cv*100, s.CritAverage(0, 0)*100)

	return sb.String()
}

func enemyMult(resist, defReduce, defIgnore float32) float32 {
	var charLvl, enemyLvl uint32
	charLvl = 90
	enemyLvl = 90

	var res float32 = .10 + resist
	switch {
	case res < 0:
		res = 1 - res/2
	case res >= .75:
		res = 1 / (4*res + 1)
	default:
		res = 1 - res
	}

	if defReduce > .9 {
		defReduce = .9
	}
	def := float32(charLvl+100) / (float32(charLvl+100) + float32(enemyLvl+100)*(1-defReduce)*(1-defIgnore))

	return def * res
}
