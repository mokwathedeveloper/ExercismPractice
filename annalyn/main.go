package main

import "fmt"

func main() {
	knightAwake := true
	archerAwake := false
	prisonerAwake := true
	dogPresent := true

	fmt.Println("Can Fast Attack:", CanFastAttack(knightAwake))
	fmt.Println("Can Spy:", CanSpy(knightAwake, archerAwake, prisonerAwake))
	fmt.Println("Can Signal Prisoner:", CanSignalPrisoner(archerAwake, prisonerAwake))
	fmt.Println("Can Free Prisoner:", CanFreePrisoner(knightAwake, archerAwake, prisonerAwake, dogPresent))
}
