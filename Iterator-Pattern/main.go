package main

func main() {
	pancakeHouseMenu := NewPancakeHouseMenu()
	dinerMenu := NewDinerMenu()

	waitress := NewWaitress(dinerMenu, pancakeHouseMenu)

	waitress.Print()
}
