package main

import "fmt"

func process(lapType string) (LapTop, error) {
	if lapType == "macBookAir" {
		return newMacBookAir(), nil
	}
	if lapType == "asusZenBook14X" {
		return newZenBook(), nil
	}

	return nil, fmt.Errorf("wrong type")
}
