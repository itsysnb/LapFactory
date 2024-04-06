package main

type ZenBook struct {
	Asus
}

func newZenBook() LapTop {
	return &ZenBook{
		Asus: Asus{
			Model:        "asusZenBook14X",
			SerialNumber: 234,
		},
	}
}
