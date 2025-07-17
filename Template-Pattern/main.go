package main

import "fmt"

func main() {
	fmt.Println("Making tea...")
	tea := NewTea()
	tea.PrepareRecipe()

	fmt.Println("Making coffee...")
	coffee := NewCoffee()
	coffee.PrepareRecipe()
}
