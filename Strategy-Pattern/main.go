package main

func main() {
	// Create different types of ducks
	mallard := NewMallardDuck()
	redhead := NewRedheadDuck()
	rubber := NewRubberDuck()
	decoy := NewDecoyDuck()

	// Display each duck
	mallard.display()
	mallard.performFly()
	mallard.performQuack()

	redhead.display()
	redhead.performFly()
	redhead.performQuack()

	rubber.display()
	rubber.performFly()
	rubber.performQuack()

	decoy.display()
	decoy.performFly()
	decoy.performQuack()
}
