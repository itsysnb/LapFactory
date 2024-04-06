package main

type Asus struct {
	Model        string
	SerialNumber int
}

func (a *Asus) Brand() string {
	return "Asus"
}

func (a *Asus) Cpu() string {
	return "intel core i5"
}

func (a *Asus) Memory() int {
	return 512
}

func (a *Asus) Price() int {
	return 600
}
