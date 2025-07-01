package main

type Beverage interface {
	getDescription() string
	setSize(string)
	getSize() string
	cost() float64
}
