package main

import (
	"math"

	"github.com/shizukayuki/ysoptimizer/pkg/good"
)

var priority = []good.CharacterKey{
	good.Keqing,
	good.Yelan,
	good.Nahida,
	good.Noelle,
	good.Furina,
	good.RaidenShogun,
	good.YaeMiko,
	good.Ganyu,
	good.KamisatoAyaka,
	good.Eula,
	good.Yoimiya,
	good.Fischl,
	good.Xiangling,
	good.Beidou,

	good.Sucrose,
	good.KukiShinobu,
	good.Nilou,
	good.SangonomiyaKokomi,

	good.Shenhe,
	good.Mona,
	good.KujouSara,
	good.YunJin,

	good.Jean,
	good.Dehya,
	good.Rosaria,
	good.Collei,
	good.Layla,
}

var config = map[good.CharacterKey]*OptimizeTarget{
	good.Ganyu: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.CryoP).
			Circlet(good.CR, good.CD, good.ATKP).
			Skip(good.HPP, good.DEFP, good.EM).
			Max(1).SlotMax(2, good.Sands, good.Goblet, good.Circlet).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.PrototypeCrescent:
				// s.Add(good.ATKP, .36)
			case good.AquaSimulacra:
				s.Add(good.HPP, .16)
				// s.AllDMG += .20
			}
			// A1: Undivided Heart
			// s.Add(good.CR, .20)
			// Elemental Resonance: Shattering Ice
			s.Add(good.CR, .15)
			return true
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch s.SetBonus {
			case good.BlizzardStrayer:
				s.Add(good.CR, .20)
				s.Add(good.CR, .20)
			}

			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.Get(good.CryoP)
			dmg *= s.CritAverage(0, 0)
			// Ice Shard DMG
			dmg *= 1.265
			return dmg
		},
	},

	good.KamisatoAyaka: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.CryoP).
			Circlet(good.CR, good.CD, good.ATKP).
			Skip(good.HPP, good.DEFP, good.EM).
			Max(1).SlotMax(2, good.Sands, good.Goblet, good.Circlet).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.MistsplitterReforged:
				s.Add(good.CryoP, .12)
				s.Add(good.CryoP, .28)
			case good.PrimordialJadeCutter:
				s.Add(good.HPP, .20)
			}
			// Elemental Resonance: Shattering Ice
			s.Add(good.CR, .15)
			return s.Get(good.ER) >= 1.40
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch t.Weapon.Key {
			case good.PrimordialJadeCutter:
				s.Add(good.ATK, s.TotalHP()*.012)
			}

			switch s.SetBonus {
			case good.BlizzardStrayer:
				s.Add(good.CR, .20)
				s.Add(good.CR, .20)
			}

			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.BurstDMG + s.Get(good.CryoP)
			dmg *= s.CritAverage(0, 0)
			// Cutting DMG
			dmg *= 1.909
			return dmg
		},
	},

	good.Rosaria: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.CryoP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP).Max(2).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.SkywardSpine:
				s.Add(good.CR, .08)
			case good.WavebreakersFin:
				v := (60 + 80 + 70*2) * .0012
				s.BurstDMG += float32(math.Min(v, .40))
			case good.TheCatch:
				s.BurstDMG += .32
				s.BurstCR += .12
			}
			// A1: Regina Probationum
			s.Add(good.CR, .12)
			// Elemental Resonance: Shattering Ice
			s.Add(good.CR, .15)
			return s.Get(good.ER) >= 1.50
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch t.Weapon.Key {
			case good.StaffOfTheScarletSands:
				s.Add(good.ATK, s.Get(good.EM)*(.52+.28*3))
			}

			switch s.SetBonus {
			case good.BlizzardStrayer:
				s.Add(good.CR, .20)
				s.Add(good.CR, .20)
			}

			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.BurstDMG + s.Get(good.CryoP)
			dmg *= s.CritAverage(s.BurstCR, 0)
			// Ice Lance DoT
			dmg *= 2.244
			return dmg
		},
	},

	good.Keqing: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.ElectroP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP, good.ER).Max(1).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.MistsplitterReforged:
				s.Add(good.ElectroP, .12)
				s.Add(good.ElectroP, .28)
			}
			// A4: Aristocratic Dignity
			// s.Add(good.CR, .15)
			// s.Add(good.ER, .15)
			return true
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.Get(good.ElectroP)
			dmg *= s.CritAverage(0, 0)
			// Slashing DMG
			dmg *= 3.024
			return dmg
		},
	},

	good.Yelan: {
		Filter: NewFilter().
			Sands(good.HPP).
			Goblet(good.HydroP).
			Circlet(good.CR, good.CD).
			Skip(good.ATKP, good.DEFP).
			Max(1).SlotMax(2, good.Sands).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.AquaSimulacra:
				s.Add(good.HPP, .16)
				s.AllDMG += .20
			}
			// A1: Turn Control
			s.Add(good.HPP, .06)
			// A4: Adapt With Ease
			s.AllDMG += .01 + .035*4
			return s.Get(good.ER) >= 1.40
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			dmg := s.TotalHP()
			dmg *= 1 + s.AllDMG + s.BurstDMG + s.Get(good.HydroP)
			dmg *= s.CritAverage(0, 0)
			// Exquisite Throw DMG
			dmg *= .088
			return dmg
		},
	},

	good.Furina: {
		Filter: NewFilter().
			Sands(good.HPP).
			Goblet(good.HydroP, good.HPP).
			Circlet(good.CR, good.CD).
			Skip(good.ATKP, good.DEFP).
			Max(1).SlotMax(2, good.Sands, good.Circlet).SlotMax(3, good.Goblet).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.PrimordialJadeCutter:
				s.Add(good.HPP, .20)
			case good.KeyOfKhajNisut:
				s.Add(good.HPP, .20)
			case good.FesteringDesire:
				s.SkillCR += .12
				s.SkillDMG += .32
			}
			return s.Get(good.ER) >= 1.60
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch s.SetBonus {
			case good.GoldenTroupe:
				s.SkillDMG += .25
			}

			hp := s.TotalHP()
			a4 := hp * .001 * .007
			if a4 > .28 {
				a4 = .28
			}

			dmg := hp
			dmg *= 1 + s.AllDMG + s.SkillDMG + a4 + s.Get(good.HydroP)
			dmg *= s.CritAverage(s.SkillCR, 0)
			// Gentilhomme Usher DMG
			dmg *= .1013
			return dmg
		},
	},

	good.Nahida: {
		Filter: NewFilter().
			Sands(good.ATKP, good.EM).
			Goblet(good.DendroP, good.EM).
			Circlet(good.CR, good.CD, good.EM).
			Skip(good.HPP, good.DEFP, good.ER).
			Max(1).SlotMax(2, good.Sands, good.Goblet, good.Circlet).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.AThousandFloatingDreams:
				s.Add(good.DendroP, .10*3)
				// s.Add(good.EM, 32*3)
			}
			return true
		},
		IgnoreEnemy: true,
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch s.SetBonus {
			case good.GildedDreams:
				s.Add(good.EM, 50*3)
			}

			em := s.Get(good.EM)
			spread := 1446.9 * (1 + (5*em)/(1200+em)) * 1.25
			if em > 200 {
				dmg := (em - 200) * .001
				if dmg > .80 {
					dmg = .80
				}
				crit := (em - 200) * .0003
				if crit > .24 {
					crit = .24
				}
				s.SkillDMG += dmg
				s.SkillCR += crit
			}

			dmg := s.TotalATK()*1.858 + em*3.715 + spread
			dmg *= 1 + s.AllDMG + s.SkillDMG + s.Get(good.DendroP)
			dmg *= s.CritAverage(s.SkillCR, 0)

			switch s.SetBonus {
			case good.DeepwoodMemories:
				dmg *= enemyMult(-.30, 0, 0)
			default:
				dmg *= enemyMult(0, 0, 0)
			}
			return dmg
		},
	},

	good.RaidenShogun: {
		Filter: NewFilter().
			Sands(good.ATKP, good.ER).
			Goblet(good.ElectroP, good.ATKP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP, good.EM).Max(1).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.PrimordialJadeWingedSpear:
				s.Add(good.ATKP, .032*2)
			case good.TheCatch:
				s.BurstDMG += .32
				s.BurstCR += .12
			}
			// Skill: Eye of Stormy Judgment
			s.BurstDMG += .0030 * 90
			return s.Get(good.ER) >= 2.00
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch t.Weapon.Key {
			case good.StaffOfTheScarletSands:
				s.Add(good.ATK, s.Get(good.EM)*(.52+.28*3))
			}

			// A4: Enlightened One
			s.Add(good.ElectroP, (s.Get(good.ER)-1)*.4)

			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.BurstDMG + s.Get(good.ElectroP)
			dmg *= s.CritAverage(s.BurstCR, 0)
			// Musou no Hitotachi DMG (40 stacks)
			dmg *= 7.214 + .07*40
			return dmg
		},
	},

	good.YaeMiko: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.ElectroP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP).Max(1).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.KagurasVerity:
				s.SkillDMG += .12 * 3
				s.Add(good.ElectroP, .12)
			}
			return s.Get(good.ER) >= 1.40
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			em := s.Get(good.EM)
			agg := 1446.9 * (1 + (5*em)/(1200+em)) * 1.15

			atk := s.TotalATK()
			electro := 1 + s.AllDMG + s.Get(good.ElectroP)

			// combo
			skill := atk * 1.706
			skill = skill*2 + (skill + agg)
			skill *= electro + s.SkillDMG + em*.0015

			burst := atk*(4.68+6.009*3) + agg*4
			burst *= electro + s.BurstDMG

			dmg := skill + burst
			dmg *= s.CritAverage(0, 0)
			return dmg
		},
	},

	good.Noelle: {
		Filter: NewFilter().
			Sands(good.ATKP, good.DEFP).
			Goblet(good.GeoP, good.DEFP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.EM).Max(1).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.SkywardPride:
				s.AllDMG += .08
			case good.SerpentSpine:
				s.AllDMG += .10 * 5
			}
			return s.Get(good.ER) >= 1.60
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch s.SetBonus {
			case good.GladiatorsFinale:
				s.NormalDMG += .35
			case good.HuskOfOpulentDreams:
				s.Add(good.GeoP, .06*4)
				s.Add(good.DEFP, .06*4)
			}

			// Burst: ATK Bonus + C6
			bonusATK := s.TotalDEF() * (.85 + .50)

			dmg := s.TotalATK() + bonusATK
			dmg *= 1 + s.AllDMG + s.NormalDMG + s.Get(good.GeoP)
			dmg *= s.CritAverage(0, 0)
			// 1-Hit DMG
			dmg *= 1.564
			return dmg
		},
	},

	good.Eula: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.PhysicalP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP, good.EM).
			Max(1).SlotMax(3, good.Sands, good.Goblet, good.Circlet).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.SkywardPride:
				s.AllDMG += .08
			case good.SerpentSpine:
				s.AllDMG += .10 * 5
			case good.BeaconOfTheReedSea:
				s.Add(good.ATKP, .20*1)
			}
			return s.Get(good.ER) >= 1.40
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch s.SetBonus {
			case good.GladiatorsFinale:
				s.NormalDMG += .35
			case good.PaleFlame:
				s.Add(good.ATKP, .09*2)
				s.Add(good.PhysicalP, .25)
			}

			phys := 1 + s.AllDMG + s.Get(good.PhysicalP)

			dmg := s.TotalATK()
			dmg *= s.CritAverage(0, 0)

			// 13 stacks combo
			var combo float32
			combo += (phys + s.NormalDMG) * (1.649 + 1.719 + 1.044*2 + 2.069)
			combo += (phys + s.BurstDMG) * (7.256 * .5)
			combo += (phys + s.BurstDMG) * (7.256 + 1.482*13)
			dmg *= combo
			return dmg
		},
	},

	good.Layla: {
		Filter: NewFilter().
			Sands(good.HPP).
			Goblet(good.HPP).
			Circlet(good.HPP).
			Skip(good.DEFP).Max(1).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.PrimordialJadeCutter:
				s.Add(good.HPP, .20)
			case good.KeyOfKhajNisut:
				s.Add(good.HPP, .20)
			}
			return s.Get(good.ER) >= 1.60 && s.SetBonus == good.TenacityOfTheMillelith
		},
		IgnoreEnemy: true,
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			return s.TotalHP()
		},
	},

	good.YunJin: {
		Filter: NewFilter().
			Sands(good.DEFP).
			Goblet(good.DEFP).
			Circlet(good.DEFP, good.CR).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.SkywardSpine:
				s.Add(good.CR, .08)
			}
			return s.Get(good.ER) >= 1.60 && s.Get(good.CR) >= .70
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			noelle, ok := optimized[good.Noelle]
			if !ok {
				return 0
			}

			if t.Character.Constellation >= 2 {
				noelle.NormalDMG += .15
			}
			if t.Character.Constellation >= 4 {
				s.Add(good.DEFP, .20)
			}

			switch s.SetBonus {
			case good.NoblesseOblige:
				s.Add(good.ATKP, .20)
				noelle.Add(good.ATKP, .20)
			case good.HuskOfOpulentDreams:
				s.Add(good.GeoP, .06*4)
				s.Add(good.DEFP, .06*4)
			}

			// Burst: DMG Increase + A4
			flatdmg := s.TotalDEF() * (.611 + .075)

			// Burst: ATK Bonus + C6
			bonusATK := noelle.TotalDEF() * (.85 + .50)

			// 1-Hit DMG
			dmg := ((noelle.TotalATK() + bonusATK) * 1.564) + flatdmg
			dmg *= 1 + noelle.AllDMG + noelle.NormalDMG + noelle.Get(good.GeoP)
			dmg *= noelle.CritAverage(0, 0)
			return dmg
		},
	},

	good.Shenhe: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.ATKP).
			Circlet(good.ATKP).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.SkywardSpine:
				s.Add(good.CR, .08)
			}
			return s.Get(good.ER) >= 2.00
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			ayaka, ok := optimized[good.KamisatoAyaka]
			if !ok {
				return 0
			}

			// A1: Deific Embrace
			ayaka.Add(good.CryoP, .15)
			// A4: Spirit Communion Seal
			ayaka.BurstDMG += .10

			switch s.SetBonus {
			case good.NoblesseOblige:
				s.Add(good.ATKP, .20)
				ayaka.Add(good.ATKP, .20)
			}

			// Icy Quill
			flatdmg := s.TotalATK() * .776

			// Cutting DMG
			dmg := (ayaka.TotalATK() * 1.909) + flatdmg
			dmg *= 1 + ayaka.AllDMG + ayaka.BurstDMG + ayaka.Get(good.CryoP)
			dmg *= ayaka.CritAverage(0, 0)
			return dmg
		},
	},

	good.Mona: {
		Filter: NewFilter().
			Sands(good.ER).
			Goblet(good.HydroP).
			Circlet(good.CR).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			return s.SetBonus == good.NoblesseOblige
		},
		IgnoreEnemy: true,
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			// A4: Waterborne Destiny
			s.Add(good.HydroP, s.Get(good.ER)*.20)

			return (1 + s.Get(good.CR)*s.Get(good.ER)) * 100
		},
	},

	good.Sucrose: {
		Filter: NewFilter().
			Sands(good.EM).
			Goblet(good.EM).
			Circlet(good.EM).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			return s.Get(good.ER) >= 1.40 && s.SetBonus == good.ViridescentVenerer
		},
		IgnoreEnemy: true,
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			return s.Get(good.EM)
		},
	},

	good.Jean: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.AnemoP).
			Circlet(good.CR, good.CD).
			Skip(good.DEFP).Max(1).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.PrimordialJadeCutter:
				s.Add(good.HPP, .20)
			case good.FesteringDesire:
				s.SkillCR += .12
				s.SkillDMG += .32
			}
			return s.Get(good.ER) >= 1.60 && s.SetBonus == good.ViridescentVenerer
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch t.Weapon.Key {
			case good.PrimordialJadeCutter:
				s.Add(good.ATK, s.TotalHP()*.012)
			}

			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.SkillDMG + s.Get(good.AnemoP)
			dmg *= s.CritAverage(s.SkillCR, 0)
			// Skill DMG
			dmg *= 4.964
			return dmg
		},
	},

	good.KukiShinobu: {
		Filter: NewFilter().
			Sands(good.EM).
			Goblet(good.EM).
			Circlet(good.EM).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			return true
		},
		IgnoreEnemy: true,
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			var bonus float32
			switch s.SetBonus {
			case good.GildedDreams:
				s.Add(good.EM, 50*3)
			case good.FlowerOfParadiseLost:
				bonus = .4 + .4*.25*4
			}

			em := s.Get(good.EM)
			hbloom := 1446.9 * (1 + (16*em)/(2000+em) + bonus) * 3
			hbloom *= enemyMult(0, 0, 1)
			return hbloom
		},
	},

	good.Nilou: {
		Filter: NewFilter().
			Sands(good.HPP).
			Goblet(good.HPP).
			Circlet(good.HPP).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.PrimordialJadeCutter:
				s.Add(good.HPP, .20)
			case good.KeyOfKhajNisut:
				s.Add(good.HPP, .20)
			}
			// A1: Court of Dancing Petals
			s.Add(good.EM, 100)
			// Elemental Resonance: Soothing Water
			s.Add(good.HPP, .25)
			return true
		},
		IgnoreEnemy: true,
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			hp := s.TotalHP()
			switch t.Weapon.Key {
			case good.KeyOfKhajNisut:
				s.Add(good.EM, hp*.0012*3)
				s.Add(good.EM, hp*.002)
			}

			// Dreamy Dance of Aeons
			a4 := hp - 30000
			if a4 < 0 {
				a4 = 0
			}
			a4 *= .001 * .09
			if a4 > 4 {
				a4 = 4
			}

			// Bloom DMG
			em := s.Get(good.EM)
			bloom := 1446.9 * (1 + (16*em)/(2000+em) + a4) * 2
			bloom *= enemyMult(0, 0, 1)
			return bloom
		},
	},

	good.SangonomiyaKokomi: {
		Filter: NewFilter().
			Sands(good.HPP).
			Goblet(good.HydroP, good.HPP).
			Circlet(good.Heal, good.HPP).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			// Passive: Flawless Strategy
			s.Add(good.CR, -1)
			s.Add(good.Heal, .25)
			return s.Get(good.ER) >= 1.60 && s.SetBonus == good.OceanHuedClam
		},
		IgnoreEnemy: true,
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			// Skill: Ripple DMG
			dmg := s.TotalATK() * 1.856
			dmg *= 1 + s.AllDMG + s.Get(good.HydroP)
			dmg *= s.CritAverage(0, 0)
			dmg *= 2

			// Skill: Regeneration
			heal := s.TotalHP()*0.0748 + 861.597
			heal *= 1 + s.Get(good.Heal)
			heal *= 2

			clam := heal
			if clam > 30000 {
				clam = 30000
			}
			clam *= .90

			clam *= enemyMult(0, 0, 1)
			dmg *= enemyMult(0, 0, 0)
			dmg += clam
			return dmg
		},
	},

	good.KujouSara: {
		Filter: NewFilter().
			Sands(good.ATKP, good.ER).
			Goblet(good.ElectroP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP, good.EM).Max(2).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.AlleyHunter:
				s.AllDMG += .30
			}
			return s.Get(good.ER) >= 1.60 && s.SetBonus == good.NoblesseOblige
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			// Noblesse Oblige
			s.Add(good.ATKP, .20)

			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.BurstDMG + s.Get(good.ElectroP)
			dmg *= s.CritAverage(0, 0)
			// Tengu Juurai: Titanbreaker DMG
			dmg *= 8.192
			return dmg
		},
	},

	good.Beidou: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.ElectroP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP).Max(2).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.SkywardPride:
				s.AllDMG += .08
			case good.SerpentSpine:
				s.AllDMG += .10 * 5
			}
			return s.Get(good.ER) >= 1.30
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.BurstDMG + s.Get(good.ElectroP)
			dmg *= s.CritAverage(0, 0)
			// Lightning DMG
			dmg *= 2.04
			return dmg
		},
	},

	good.Yoimiya: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.PyroP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP, good.ER).
			Max(1).SlotMax(2, good.Sands, good.Goblet, good.Circlet).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.Rust:
				s.NormalDMG += .80
				s.ChargedDMG += -.10
			}
			// A1: Tricks of the Trouble-Maker
			s.Add(good.PyroP, .02*6)
			return true
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			em := s.Get(good.EM)
			amp := 1 + (2.778*em)/(1400+em)

			switch s.SetBonus {
			case good.ShimenawasReminiscence:
				s.NormalDMG += .50
			case good.CrimsonWitchOfFlames:
				amp += .15
				s.Add(good.PyroP, .15*.5)
			}

			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.NormalDMG + s.Get(good.PyroP)
			dmg *= s.CritAverage(0, 0)
			// Skill: Normal Attack DMG
			dmg *= 1.588

			// na string (vape)
			vape := dmg * (.599 + 1.495 + 1.78) * amp * 1.5
			novape := dmg * (.599 + 1.150 + .781*2)
			dmg = vape + novape
			return dmg
		},
	},

	good.Fischl: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.ElectroP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP, good.ER).
			Max(2).SlotMax(3, good.Sands, good.Goblet, good.Circlet).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.Rust:
				s.NormalDMG += .80
				s.ChargedDMG += -.10
			case good.TheStringless:
				s.SkillDMG += .48
				s.BurstDMG += .48
			}
			return true
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			em := s.Get(good.EM)
			agg := 1446.9 * (1 + (5*em)/(1200+em)) * 1.15

			switch s.SetBonus {
			case good.GoldenTroupe:
				s.SkillDMG += .25
			}

			// Oz's ATK DMG
			dmg := s.TotalATK() * 1.776
			dmg = dmg*2 + (dmg + agg)
			dmg *= 1 + s.AllDMG + s.SkillDMG + s.Get(good.ElectroP)
			dmg *= s.CritAverage(s.SkillCR, 0)
			return dmg
		},
	},

	good.Xiangling: {
		Filter: NewFilter().
			Sands(good.ATKP, good.ER).
			Goblet(good.PyroP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP).
			Max(1).SlotMax(2, good.Sands, good.Goblet, good.Circlet).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.TheCatch:
				s.BurstCR += .12
				s.BurstDMG += .32
			}
			return s.Get(good.ER) >= 2.00
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			em := s.Get(good.EM)
			amp := 1 + (2.778*em)/(1400+em)
			switch s.SetBonus {
			case good.CrimsonWitchOfFlames:
				amp += .15
				s.Add(good.PyroP, .15*.5)
			}

			dmg := s.TotalATK()
			dmg *= 1 + s.AllDMG + s.BurstDMG + s.Get(good.PyroP)
			dmg *= s.CritAverage(s.BurstCR, 0)
			// Pyronado DMG
			novape := dmg * 2.24
			vape := novape * amp * 1.5
			dmg = vape + novape
			return dmg
		},
	},

	good.Dehya: {
		Filter: NewFilter().
			Sands(good.ATKP, good.HPP).
			Goblet(good.PyroP).
			Circlet(good.CR, good.CD).
			Skip(good.DEFP).Max(1).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.SkywardPride:
				s.AllDMG += .08
			case good.BeaconOfTheReedSea:
				s.Add(good.ATKP, .20*2)
				s.Add(good.HPP, .32)
			}
			if t.Character.Constellation >= 1 {
				s.Add(good.HPP, .20)
			}
			return true
		},
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			hp := s.TotalHP()
			flatdmg := hp * .0406
			if t.Character.Constellation >= 1 {
				flatdmg += hp * .06
			}

			// Incineration Drive DMG
			dmg := (s.TotalATK() * 3.024) + flatdmg
			dmg *= 1 + s.AllDMG + s.BurstDMG + s.Get(good.PyroP)
			dmg *= s.CritAverage(0, 0)
			return dmg
		},
	},

	good.Collei: {
		Filter: NewFilter().
			Sands(good.ATKP).
			Goblet(good.DendroP).
			Circlet(good.CR, good.CD).
			Skip(good.HPP, good.DEFP).Max(2).
			Build(),
		Buffs: func(t *OptimizeTarget, s *OptimizeState) bool {
			switch t.Weapon.Key {
			case good.TheStringless:
				s.SkillDMG += .48
				s.BurstDMG += .48
			}
			return true
		},
		IgnoreEnemy: true,
		Target: func(t *OptimizeTarget, s *OptimizeState) float32 {
			switch s.SetBonus {
			case good.GildedDreams:
				s.Add(good.EM, 50*2)
			}

			em := s.Get(good.EM)
			spread := 1446.9 * (1 + (5*em)/(1200+em)) * 1.25

			// Skill DMG
			dmg := (s.TotalATK() * 2.570) + spread
			dmg *= 1 + s.AllDMG + s.SkillDMG + s.Get(good.DendroP)
			dmg *= s.CritAverage(s.SkillCR, 0)

			switch s.SetBonus {
			case good.DeepwoodMemories:
				dmg *= enemyMult(-.30, 0, 0)
			default:
				dmg *= enemyMult(0, 0, 0)
			}
			return dmg
		},
	},
}
