package main

type CaffeineBeverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
	CustomerWantsCondiments() bool
	PrepareRecipe()
}


