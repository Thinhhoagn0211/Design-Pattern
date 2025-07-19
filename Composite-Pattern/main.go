package main

func main() {
	pancakeHouseMenu := NewMenu("PANCAKE HOUSE MENU", "Breakfast")
	dinerMenu := NewMenu("DINER MENU", "Lunch")
	cafeMenu := NewMenu("CAFE MENU", "Dinner")
	desertMenu := NewMenu("DESSERT MENU", "Desserts")

	allMenus := NewMenu("ALL MENUS", "All menus combined")
	allMenus.Add(pancakeHouseMenu)
	allMenus.Add(dinerMenu)
	allMenus.Add(cafeMenu)

	dinerMenu.Add(NewMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", 2.99, true))
	dinerMenu.Add(desertMenu)

	desertMenu.Add(NewMenuItem("Apple Pie", "Apple pie with a flakey crust, topped with vanilla ice cream", 1.59, true))

	waitress := NewWaitress(allMenus)
	println(waitress.PrintMenu()) // Print the combined menu
}

