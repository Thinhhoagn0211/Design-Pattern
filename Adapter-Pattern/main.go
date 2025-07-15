package main

func main() {

	// Create a new duck
	duck := NewMallardDuck()

	// Make the duck quack and fly
	duck.Quack()
	duck.Fly()

	// Create a turkey
	turkey := NewWildTurkey()

	// Create an adapter for the turkey
	turkeyAdapter := NewTurkeyAdapter(turkey)

	// Make the turkey adapter quack and fly
	turkeyAdapter.Quack()
	turkeyAdapter.Fly()
}
