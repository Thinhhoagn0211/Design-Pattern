package main

import (
	"fmt"
	"sync"
)



type ChocolateBoiler struct {
	empty   bool
	boiled  bool
}

var instance *ChocolateBoiler
var once sync.Once
func NewChocolateBoiler() *ChocolateBoiler {
	return &ChocolateBoiler{
		empty:  true,
		boiled: false,
	}
}

func (b *ChocolateBoiler) IsEmpty() bool {
	return b.empty
}

func (b *ChocolateBoiler) IsBoiled() bool {
	return b.boiled
}

func (b *ChocolateBoiler) Fill() {
	if b.IsEmpty() {
		b.empty = false
		b.boiled = false
		// Fill the boiler with a milk/chocolate mixture
	}
}

func (b *ChocolateBoiler) Drain() {
	if !b.IsEmpty() && b.IsBoiled() {
		// Drain the boiled milk/chocolate mixture
		b.empty = true
	}
}

func (b *ChocolateBoiler) Boil() {
	if !b.IsEmpty() && !b.IsBoiled() {
		// Bring the contents to a boil
		b.boiled = true
	}
}

func GetChocolateBoilerInstance() *ChocolateBoiler {
	once.Do(func() {
		instance = NewChocolateBoiler()
		fmt.Println("Creating new ChocolateBoiler instance")
	})
	return instance
}

