package main

import "fmt"

type Beverage struct {}

func (b *Beverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (b *Beverage) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (b *Beverage) CustomerWantsCondiments() bool {
	// Default implementation, can be overridden
	return true
}
