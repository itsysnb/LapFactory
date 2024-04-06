package main

type Mac struct {
	Model        string
	SerialNumber int
}

func (m *Mac) Brand() string {
	return "Apple"
}

func (m *Mac) Cpu() string {
	return "air"
}

func (m *Mac) Memory() int {
	return 512
}

func (m *Mac) Price() int {
	return 590
}
