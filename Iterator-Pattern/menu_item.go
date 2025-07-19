package main

type MenuItem struct {
	Name        string
	Description string
	Price       float64
	Vegetarian  bool
}

func NewMenuItem(name, description string, price float64, vegetarian bool) *MenuItem {
	return &MenuItem{
		Name:        name,
		Description: description,
		Price:       price,
		Vegetarian:  vegetarian,
	}
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

