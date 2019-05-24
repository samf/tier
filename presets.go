package tier

var (
	// Bytes is for bytes, kilobytes, etc.
	Bytes = MakeTiered(
		Tier{
			Name:   "bytes",
			Abbrev: "b",
			Units:  1,
		},
		Tier{
			Name:   "kilobytes",
			Abbrev: "k",
			Units:  1024,
		},
		Tier{
			Name:   "megabytes",
			Abbrev: "m",
			Units:  1024 * 1024,
		},
		Tier{
			Name:   "gigabytes",
			Abbrev: "g",
			Units:  1024 * 1024 * 1024,
		},
		Tier{
			Name:   "terabytes",
			Abbrev: "t",
			Units:  1024 * 1024 * 1024 * 1024,
		},
		Tier{
			Name:   "petabytes",
			Abbrev: "p",
			Units:  1024 * 1024 * 1024 * 1024 * 1024,
		},
		Tier{
			Name:   "exabytes",
			Abbrev: "e",
			Units:  1024 * 1024 * 1024 * 1024 * 1024 * 1024,
		},
	)

	// Time is for weeks, days, hours, etc.
	Time = MakeTiered(
		Tier{
			Name:   "seconds",
			Abbrev: "s",
			Units:  1,
		},
		Tier{
			Name:   "minutes",
			Abbrev: "m",
			Units:  60,
		},
		Tier{
			Name:   "hours",
			Abbrev: "h",
			Units:  60 * 60,
		},
		Tier{
			Name:   "days",
			Abbrev: "d",
			Units:  60 * 60 * 24,
		},
		Tier{
			Name:   "weeks",
			Abbrev: "w",
			Units:  60 * 60 * 24 * 7,
		},
	)
)
