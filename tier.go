package tierflat

import (
	"fmt"
	"sort"
)

// Tier is a single tier to be used in a Tiered type.
type Tier struct {
	Name   string
	Units  int
	Abbrev string
	Amount int
}

// Tiered is a type for representing a system of tiers. It should be created via the
// MakeTiered call.
type Tiered []Tier

// MakeTiered creates a well-ordered Tiered out of a list of Tiers.
func MakeTiered(tiers ...Tier) Tiered {
	sort.Slice(tiers, func(i, j int) bool {
		return tiers[j].Units < tiers[i].Units
	})

	return tiers
}

// FlatToTiered takes a tiered as a prototype, and converts a flat amount
// into a Tiered.
func (t Tiered) FlatToTiered(amount int) Tiered {
	tiers := MakeTiered(t...)

	for i := range tiers {
		if amount < tiers[i].Units {
			continue
		}

		tiers[i].Amount = amount / tiers[i].Units
		amount %= tiers[i].Units
	}

	return tiers
}

func (t Tiered) String() string {
	var out string

	for _, ti := range t {
		if ti.Amount != 0 {
			out += fmt.Sprintf("%v%v", ti.Amount, ti.Abbrev)
		}
	}
	if out == "" {
		last := len(t) - 1
		out = fmt.Sprintf("0%v", t[last].Abbrev)
	}

	return out
}
