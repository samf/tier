package tierflat

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
)
