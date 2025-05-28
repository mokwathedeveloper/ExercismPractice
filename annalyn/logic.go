package main

// CanFastAttack returns true if the knight is asleep
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy returns true if at least one of knight, archer or prisoner is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner returns true if the prisoner is awake and the archer is asleep
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

// CanFreePrisoner returns true if Annalyn can free the prisoner based on the given conditions
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent {
		return !archerIsAwake
	}
	return prisonerIsAwake && !knightIsAwake && !archerIsAwake
}
