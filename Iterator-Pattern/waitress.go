package main

import "fmt"

type Waitress struct {
	dinerMenu *DinerMenu
	pancakeHouseMenu *PancakeHouseMenu
}

func NewWaitress(dinerMenu *DinerMenu, pancakeHouseMenu *PancakeHouseMenu) *Waitress {
	return &Waitress{
		dinerMenu: dinerMenu,
		pancakeHouseMenu: pancakeHouseMenu,
	}
}

func (w *Waitress) Print() {
	pancakeIterator := w.pancakeHouseMenu.CreateIterator()
	dinerIterator := w.dinerMenu.CreateIterator()

	fmt.Println("MENU\n----\nBREAKFAST")
	w.PrintMenu(pancakeIterator)
	fmt.Println("\nLUNCH")
	w.PrintMenu(dinerIterator)
}

func (w *Waitress) PrintMenu(iterator Iterator) {
	for iterator.HasNext() {
		menuItem := iterator.Next()
		fmt.Printf("  %s, %.2f -- %s\n", menuItem.GetName(), menuItem.GetPrice(), menuItem.GetDescription())
	}
}
