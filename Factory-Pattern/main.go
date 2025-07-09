package main

import "fmt"

func main() {
	nyStore := NewNYPizzaStore()
	chicagoStore := NewChicagoPizzaStore()

	fmt.Println("Ordering a NY Style Cheese Pizza:")
	nyPizza := nyStore.OrderPizza("cheese")
	fmt.Printf("Ordered: %s\n", nyPizza.GetName())

	fmt.Println("\nOrdering a Chicago Style Pepperoni Pizza:")
	chicagoPizza := chicagoStore.OrderPizza("pepperoni")
	fmt.Printf("Ordered: %s\n", chicagoPizza.GetName())
} 
