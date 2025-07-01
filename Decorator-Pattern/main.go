package main

import "fmt"

func main() {
	// Create a beverage with no condiments
	beverage := NewEspresso()
	fmt.Println(beverage.getDescription() + " $" + fmt.Sprintf("%.2f", beverage.cost()))

	beverage2 := NewDarkRoast()
	beverage2.setSize("Large")
	beverage2 = NewMocha(beverage2)
	beverage2 = NewMocha(beverage2)
	beverage2 = NewWhip(beverage2)
	fmt.Println(beverage2.getDescription() + " $" + fmt.Sprintf("%.2f", beverage2.cost()))

	beverage3 := NewHouseBlend()
	beverage3.setSize("Small")
	beverage3 = NewMocha(beverage3)
	beverage3 = NewWhip(beverage3)
	fmt.Println(beverage3.getDescription() + " $" + fmt.Sprintf("%.2f", beverage3.cost()))
}
