package main

import "fmt"

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) swim() {
	fmt.Println("All ducks float, even decoys!")
}

func (d *Duck) display() {
	// This method should be implemented by concrete duck types
	fmt.Println("Displaying duck")
}

func (d *Duck) performQuack() {
	d.quackBehavior.quack()
}

func (d *Duck) performFly() {
	d.flyBehavior.fly()
}

func (d *Duck) setFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

func (d *Duck) setQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}
