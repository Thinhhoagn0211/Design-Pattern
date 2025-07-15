package main

import "fmt"

type MallardDuck struct {}

func NewMallardDuck() Duck {
	return &MallardDuck{}
}

func (d *MallardDuck) Quack() {
	fmt.Println("Quack")
}

func (d *MallardDuck) Fly() {
	fmt.Println("I'm flying")
}

type WildTurkey struct {}

func NewWildTurkey() Turkey {
	return &WildTurkey{}
}

func (t *WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}

func (t *WildTurkey) Fly() {
	fmt.Println("I'm flying a short distance")
}

