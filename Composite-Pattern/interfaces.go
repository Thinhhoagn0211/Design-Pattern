package main

type MenuComponent interface {
	Add(component MenuComponent)
	Remove(component MenuComponent)
	GetChild(index int) MenuComponent
	GetName() string
	GetDescription() string
	GetPrice() float64
	IsVegetarian() bool
	Print() string
}
