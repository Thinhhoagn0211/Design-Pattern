package main

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type PizzaStore interface {
	OrderPizza(pizzaType string) Pizza
	//CreatePizza(pizzaType string) Pizza
}
