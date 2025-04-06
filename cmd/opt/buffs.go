package main

import "github.com/shizukayuki/ysoptimizer/pkg/good"

func buffs(t *OptimizeTarget, s *OptimizeState) bool {
	r := float32(t.Weapon.Refinement - 1)
	switch t.Weapon.Key {
	case good.FesteringDesire:
		s.SkillDMG += .16 + 0.04*r
		s.SkillCR += .06 + .015*r
	case good.KeyOfKhajNisut:
		s.Add(good.HPP, .20+.05*r)
	case good.MistsplitterReforged:
		var dmg float32 = .12 + .03*r
		s.Add(good.PyroP, dmg)
		s.Add(good.HydroP, dmg)
		s.Add(good.DendroP, dmg)
		s.Add(good.ElectroP, dmg)
		s.Add(good.AnemoP, dmg)
		s.Add(good.CryoP, dmg)
		s.Add(good.GeoP, dmg)
	case good.PrimordialJadeCutter:
		s.Add(good.HPP, .20+.05*r)
		s.Add(good.ATK, s.TotalHP()*(.012+.003*r))
	case good.UrakuMisugiri:
		s.NormalDMG += .16 + .04*r
		s.SkillDMG += .24 + .06*r
		s.Add(good.DEFP, .20+.05*r)

	case good.AThousandBlazingSuns:
		s.Add(good.CD, .20+.05*r)
		s.Add(good.ATKP, .28+.07*r)
	case good.BeaconOfTheReedSea:
		s.Add(good.ATKP, .20+.05*r)
		s.Add(good.HPP, .32+.08*r)
	case good.RedhornStonethresher:
		s.Add(good.DEFP, .28+.07*r)
	case good.SerpentSpine:
		s.AllDMG += 5 * (.06 + .01*r)
	case good.SkywardPride:
		s.AllDMG += .08 + .02*r
	case good.Verdict:
		s.Add(good.ATKP, .20+.05*r)

	case good.AquaSimulacra:
		s.Add(good.HPP, .16+.04*r)
	case good.Rust:
		s.NormalDMG += .40 + .10*r
		s.ChargedDMG += -.10
	case good.TheStringless:
		s.SkillDMG += .24 + .06*r
		s.BurstDMG += .24 + .06*r

	case good.PrimordialJadeWingedSpear:
		s.Add(good.ATKP, 2*(.032+.007*r))
	case good.SkywardSpine:
		s.Add(good.CR, .08+.02*r)
	case good.StaffOfTheScarletSands:
		s.Add(good.ATK, s.Get(good.EM)*(.52+.13*r))
	case good.TheCatch:
		s.BurstDMG += .16 + .04*r
		s.BurstCR += .06 + .015*r
	case good.WavebreakersFin:
		s.BurstDMG += min(4*70*(.0012+.0003*r), .40+.10*r)

	case good.StarcallersWatch:
		s.Add(good.EM, 100+25*r)
	}
	return t.Buffs(t, s)
}
