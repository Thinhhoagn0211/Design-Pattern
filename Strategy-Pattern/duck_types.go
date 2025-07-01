package main

type MallardDuck struct{}

func NewMallardDuck() *Duck {
	duck := &Duck{}
	duck.setFlyBehavior(&FlyWithWings{})
	duck.setQuackBehavior(&Quack{})
	return duck
}

func (d *MallardDuck) display() {
	println("I'm a Mallard Duck")
}

type RedheadDuck struct{}

func NewRedheadDuck() *Duck {
	duck := &Duck{}
	duck.setFlyBehavior(&FlyWithWings{})
	duck.setQuackBehavior(&Quack{})
	return duck
}

func (d *RedheadDuck) display() {
	println("I'm a Redhead Duck")
}

type RubberDuck struct{}

func NewRubberDuck() *Duck {
	duck := &Duck{}
	duck.setFlyBehavior(&FlyNoWay{})
	duck.setQuackBehavior(&Squeak{})
	return duck
}

func (d *RubberDuck) display() {
	println("I'm a Rubber Duck")
}

type DecoyDuck struct{}

func NewDecoyDuck() *Duck {
	duck := &Duck{}
	duck.setFlyBehavior(&FlyNoWay{})
	duck.setQuackBehavior(&MuteQuack{})
	return duck
}
func (d *DecoyDuck) display() {
	println("I'm a Decoy Duck")
}
