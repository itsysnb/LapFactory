package main

import "fmt"

func main() {
	macBook, _ := process("macBookAir")
	zenBook, _ := process("asusZenBook14X")

	printDetails(macBook)
	printDetails(zenBook)
}

func printDetails(l LapTop) {
	fmt.Printf("brand: %s", l.Brand())
	fmt.Printf("cpu: %s", l.Cpu())
}
