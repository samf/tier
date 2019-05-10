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
type Tiered struct {
	Tiers []Tier

	sorted bool
}

// MakeTiered creates a well-ordered Tiered out of a list of Tiers.
func MakeTiered(tiers ...Tier) Tiered {
	t := Tiered{
		Tiers: tiers,
	}

	t.sort()

	return t
}

// FlatToTiered takes a tiered as a prototype, and converts a flat amount
// into a Tiered.
func (t Tiered) FlatToTiered(amount int) Tiered {
	if !t.sorted {
		t.sort()
	}

	for i := range t.Tiers {
		if amount < t.Tiers[i].Units {
			continue
		}

		t.Tiers[i].Amount = amount / t.Tiers[i].Units
		amount %= t.Tiers[i].Units
	}

	return t
}

func (t Tiered) String() string {
	var out string

	if !t.sorted {
		t.sort()
	}

	for _, ti := range t.Tiers {
		if ti.Amount != 0 {
			out += fmt.Sprintf("%v%v", ti.Amount, ti.Abbrev)
		}
	}
	if out == "" {
		last := len(t.Tiers) - 1
		out = fmt.Sprintf("0%v", t.Tiers[last].Abbrev)
	}

	return out
}

func (t *Tiered) sort() {
	sort.Slice(t.Tiers, func(i, j int) bool {
		return t.Tiers[j].Units < t.Tiers[i].Units
	})

	t.sorted = true
}
