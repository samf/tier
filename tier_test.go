package tierflat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	timeish := MakeTiered(
		Tier{
			Units:  1,
			Abbrev: "s",
		},
		Tier{
			Units:  60,
			Abbrev: "m",
		},
		Tier{
			Units:  3600,
			Abbrev: "h",
		},
	)
	english := MakeTiered(
		Tier{
			Units:  1,
			Abbrev: "\"",
		},
		Tier{
			Units:  12,
			Abbrev: "'",
		},
	)

	t.Run("strings", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal(`6'3"`, english.FlatToTiered(75).String())
		assert.Equal("1h1s", timeish.FlatToTiered(3601).String())
		assert.Equal("2h2m2s", timeish.FlatToTiered(7322).String())
	})
}
