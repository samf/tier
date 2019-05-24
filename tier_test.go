package tier

import (
	"fmt"
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

		assert.Equal(`6'3"`, english.Make(75).String())
		assert.Equal(`4"`, english.Make(4).String())
		assert.Equal("1h1s", timeish.Make(3601).String())
		assert.Equal("2h2m2s", timeish.Make(7322).String())
		assert.Equal("8h1s", timeish.Make(28801).String())

		assert.Equal(`7'4"`, fmt.Sprintf("%v", english.Make(88)))

		for i := range english.tiers {
			assert.Zero(english.tiers[i].Amount)
		}
		for i := range timeish.tiers {
			assert.Zero(timeish.tiers[i].Amount)
		}
	})
	t.Run("values", func(t *testing.T) {
		assert := assert.New(t)

		cleese := english.Make(77)
		assert.Equal(`6'5"`, cleese.String())
		assert.Equal(int64(77), cleese.Value())
	})
	t.Run("presets", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal("32k", Bytes.Make(32768).String())
	})
	t.Run("shorten", func(t *testing.T) {
		assert := assert.New(t)

		blob := Bytes.Make(33333)
		assert.Equal("32k", blob.Short())
		assert.NotEqual("32k", blob.String())
	})
}
