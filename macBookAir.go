package main

type MacBookAir struct {
	Mac
}

func newMacBookAir() LapTop {
	return &MacBookAir{
		Mac: Mac{
			Model:        "macBookAir",
			SerialNumber: 435,
		},
	}
}
