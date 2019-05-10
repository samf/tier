package tierflat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	timeish := MakeTiered(
		Tier{
			Name:   "seconds",
			Units:  1,
			Abbrev: "s",
		},
		Tier{
			Name:   "minutes",
			Units:  60,
			Abbrev: "m",
		},
		Tier{
			Name:   "hours",
			Units:  3600,
			Abbrev: "h",
		},
	)
	english := MakeTiered(
		Tier{
			Name:   "inches",
			Units:  1,
			Abbrev: "\"",
		},
		Tier{
			Name:   "feet",
			Units:  12,
			Abbrev: "'",
		},
	)

	t.Run("strings", func(t *testing.T) {
		assert := assert.New(t)

		for i := range english.Tiers {
			assert.Zero(english.Tiers[i].Amount)
		}

		assert.Equal(`6'3"`, english.From(75).String())
		assert.Equal(`4"`, english.From(4).String())
		assert.Equal("1h1s", timeish.From(3601).String())
		assert.Equal("2h2m2s", timeish.From(7322).String())
		assert.Equal("8h1s", timeish.From(28801).String())

		for i := range english.Tiers {
			assert.Zero(english.Tiers[i].Amount)
		}
		for i := range timeish.Tiers {
			assert.Zero(timeish.Tiers[i].Amount)
		}
	})
}
