package main

import (
	"slices"

	"github.com/shizukayuki/ysoptimizer/pkg/good"
)

type Filter struct {
	sands   []good.StatKey
	goblet  []good.StatKey
	circlet []good.StatKey

	max  map[good.SlotKey]int
	skip []good.StatKey
	slow *bool
}

func NewFilter() Filter {
	return Filter{
		max:  map[good.SlotKey]int{},
		slow: slow,
	}.Max(1)
}

func (f Filter) Sands(v ...good.StatKey) Filter {
	f.sands = v
	return f
}

func (f Filter) Goblet(v ...good.StatKey) Filter {
	f.goblet = v
	return f
}

func (f Filter) Circlet(v ...good.StatKey) Filter {
	f.circlet = v
	return f
}

func (f Filter) Skip(v ...good.StatKey) Filter {
	f.skip = v
	return f
}

func (f Filter) Max(v int) Filter {
	for _, slot := range good.SlotKeyValues() {
		f.max[slot] = v
	}
	return f
}

func (f Filter) SlotMax(v int, slots ...good.SlotKey) Filter {
	for _, slot := range slots {
		f.max[slot] = v
	}
	return f
}

func (f Filter) Build() ArtifactFilter {
	for i := 0; i < len(f.skip); i++ {
		switch f.skip[i] {
		case good.HPP:
			f.skip = append(f.skip, good.HP)
		case good.ATKP:
			f.skip = append(f.skip, good.ATK)
		case good.DEFP:
			f.skip = append(f.skip, good.DEF)
		}
	}

	shouldSkip := func(a *good.Artifact) bool {
		if *f.slow {
			return false
		}

		c := 0
		for _, s := range a.Substats {
			if slices.Contains(f.skip, s.Key) {
				c++
			}
		}
		return c > f.max[a.SlotKey]
	}

	return func(a *good.Artifact) bool {
		switch a.SlotKey {
		case good.Flower:
			return !shouldSkip(a)
		case good.Plume:
			return !shouldSkip(a)
		case good.Sands:
			return slices.Contains(f.sands, a.MainStatKey) && !shouldSkip(a)
		case good.Goblet:
			return slices.Contains(f.goblet, a.MainStatKey) && !shouldSkip(a)
		case good.Circlet:
			return slices.Contains(f.circlet, a.MainStatKey) && !shouldSkip(a)
		}
		return false
	}
}
