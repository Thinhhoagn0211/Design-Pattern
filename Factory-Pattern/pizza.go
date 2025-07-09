package main

import "fmt"

type BasePizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (p *BasePizza) Prepare() {
	fmt.Printf("Preparing %s\n", p.name)
	fmt.Printf("Tossing %s...\n", p.dough)
	fmt.Printf("Adding sauce: %s\n", p.sauce)
	for _, topping := range p.toppings {
		fmt.Printf("Adding topping: %s\n", topping)
	}
}

func (p *BasePizza) Bake() {
	fmt.Printf("Baking %s\n", p.name)
}

func (p *BasePizza) Cut() {
	fmt.Printf("Cutting %s\n", p.name)
}

func (p *BasePizza) Box() {
	fmt.Printf("Boxing %s\n", p.name)
}

func (p *BasePizza) GetName() string {
	return p.name
}

type NYStyleCheesePizza struct{
	BasePizza
}

func NewNYStyleCheesePizza() Pizza {
	pizza := &NYStyleCheesePizza{}
	pizza.name = "NY Style Cheese Pizza"
	pizza.dough = "Thin Crust Dough"
	pizza.sauce = "Marinara Sauce"
	pizza.toppings = []string{"Grated Reggiano Cheese"}
	return pizza
}

type NYStylePepperoniPizza struct {
	BasePizza
}

func NewNYStylePepperoniPizza() Pizza {
	pizza := &NYStylePepperoniPizza{}
	pizza.name = "NY Style Pepperoni Pizza"
	pizza.dough = "Thin Crust Dough"
	pizza.sauce = "Marinara Sauce"
	pizza.toppings = []string{"Sliced Pepperoni", "Grated Reggiano Cheese"}
	return pizza
}

type NYStyleVeggiePizza struct {
	BasePizza
}
func NewNYStyleVeggiePizza() Pizza {
	pizza := &NYStyleVeggiePizza{}
	pizza.name = "NY Style Veggie Pizza"
	pizza.dough = "Thin Crust Dough"
	pizza.sauce = "Marinara Sauce"
	pizza.toppings = []string{"Shredded Mozzarella Cheese", "Diced Onion", "Sliced Mushrooms", "Sliced Red Pepper"}
	return pizza
}

type ChicagoStyleCheesePizza struct {
	BasePizza
}

func NewChicagoStyleCheesePizza() Pizza {
	pizza := &ChicagoStyleCheesePizza{}
	pizza.name = "Chicago Style Cheese Pizza"
	pizza.dough = "Extra Thick Crust Dough"
	pizza.sauce = "Plum Tomato Sauce"
	pizza.toppings = []string{"Shredded Mozzarella Cheese"}
	return pizza
}

type ChicagoStylePepperoniPizza struct {
	BasePizza
}

func NewChicagoStylePepperoniPizza() Pizza {
	pizza := &ChicagoStylePepperoniPizza{}
	pizza.name = "Chicago Style Pepperoni Pizza"
	pizza.dough = "Extra Thick Crust Dough"
	pizza.sauce = "Plum Tomato Sauce"
	pizza.toppings = []string{"Sliced Pepperoni", "Shredded Mozzarella Cheese"}
	return pizza
}

type ChicagoStyleVeggiePizza struct {
	BasePizza
}
func NewChicagoStyleVeggiePizza() Pizza {
	pizza := &ChicagoStyleVeggiePizza{}
	pizza.name = "Chicago Style Veggie Pizza"
	pizza.dough = "Extra Thick Crust Dough"
	pizza.sauce = "Plum Tomato Sauce"
	pizza.toppings = []string{"Shredded Mozzarella Cheese", "Diced Onion", "Sliced Mushrooms", "Sliced Red Pepper"}
	return pizza
}




