package main

import "fmt"

type Tea struct {
	Beverage
}

func NewTea() CaffeineBeverage {
	return &Tea{
		Beverage: Beverage{},
	}
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}

func (t *Tea) CustomerWantsCondiments() bool {
	// Here we can simulate user input or a condition
	return false
}

func (t *Tea) PrepareRecipe() {
	t.Beverage.BoilWater()
	t.Brew()
	t.Beverage.PourInCup()
	if t.CustomerWantsCondiments() {
		t.AddCondiments()
	}
}

type Coffee struct {
	Beverage
}

func NewCoffee() CaffeineBeverage {
	return &Coffee{
		Beverage: Beverage{},
	}
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk")
}

func (c *Coffee) CustomerWantsCondiments() bool {
	// Here we can simulate user input or a condition
	return true // Assume user wants CustomerWantsCondiments
}

func (c *Coffee) PrepareRecipe() {
	c.Beverage.BoilWater()
	c.Brew()
	c.Beverage.PourInCup()
	if c.CustomerWantsCondiments() {
		c.AddCondiments()
	}
}


