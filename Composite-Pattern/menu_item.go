package main

import "fmt"

type MenuItem struct {
	Name        string
	Description string
	Price       float64
	Vegetarian  bool
}

func NewMenuItem(name, description string, price float64, vegetarian bool) MenuComponent {
	return &MenuItem{
		Name:        name,
		Description: description,
		Price:       price,
		Vegetarian:  vegetarian,
	}
}

func (m *MenuItem) Add(component MenuComponent) {
	// MenuItems cannot have children, so this method does nothing
} 

func (m *MenuItem) Remove(component MenuComponent) {
	// MenuItems cannot have children, so this method does nothing
}

func (m *MenuItem) GetChild(index int) MenuComponent {
	// MenuItems do not have children, so we return nil
	return nil
}

func (m *MenuItem) GetName() string {
	return m.Name
}

func (m *MenuItem) GetDescription() string {
	return m.Description
}

func (m *MenuItem) GetPrice() float64 {
	return m.Price
}

func (m *MenuItem) IsVegetarian() bool {
	return m.Vegetarian
}

func (m *MenuItem) Print() string {
	vegetarianStr := ""
	if m.Vegetarian {
		vegetarianStr = " (V)"
	}
	return m.Name + vegetarianStr + ", " + m.Description + " -- $" + fmt.Sprintf("%.2f", m.Price)
}

