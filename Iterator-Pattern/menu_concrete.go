package main

import "fmt"


const maxItems = 6
type DinerMenu struct {
	items [maxItems]*MenuItem
	maxItems int
	numberOfItems int
}

func NewDinerMenu() *DinerMenu {
	menuItems := &DinerMenu{}

	menuItems.AddItem("Vegetarian BLT","Fakin' Bacon with lettuce & tomato on whole wheat", 2.99, true)
	menuItems.AddItem("BLT", "Bacon with lettuce & tomato on whole wheat", 2.99, false)
	menuItems.AddItem("Soup of the day", "A bowl of the soup of the day, with a side of potato salad", 3.29, false)
	menuItems.AddItem("Hotdog", "A hot dog, with sauerkraut, relish, onions, topped with cheese", 3.05, false)
	menuItems.AddItem("Steamed Veggies and Brown Rice", "Steamed vegetables over brown rice", 3.99, true)
	menuItems.AddItem("Pasta", "Spaghetti with marinara sauce, and a slice of sourdough bread", 3.89, true)
	return menuItems
}

func (d *DinerMenu) CreateIterator() Iterator {
	return NewDinerMenuIterator(d.items)
}

func (d *DinerMenu) AddItem(name, description string, price float64, vegetarian bool) {
	if d.numberOfItems >= maxItems {
		fmt.Println("Sorry, menu is full! Can't add item to menu.")
		return
	}
	item := NewMenuItem(name, description, price, vegetarian)
	d.items[d.numberOfItems] = item
	d.numberOfItems++
}

func (d *DinerMenu) GetMenuItems() [maxItems]*MenuItem {
	return d.items
}

type PancakeHouseMenu struct {
	items []*MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	menuItems := &PancakeHouseMenu{
		items: make([]*MenuItem, 0),
	}

	menuItems.AddItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", 2.99, true)
	menuItems.AddItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", 2.99, false)
	menuItems.AddItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", 3.49, true)
	menuItems.AddItem("Waffles", "Waffles with your choice of blueberries or strawberries", 3.59, true)

	return menuItems	
}

func (p *PancakeHouseMenu) CreateIterator() Iterator {
	return NewPancakeHouseMenuIterator(p.items)
}

func (p *PancakeHouseMenu) AddItem(name, description string, price float64, vegetarian bool) {
	item := NewMenuItem(name, description, price, vegetarian)
	p.items = append(p.items, item)
}

func (p *PancakeHouseMenu) GetMenuItems() []*MenuItem {
	return p.items
}


