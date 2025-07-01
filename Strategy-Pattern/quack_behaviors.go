package main

import "fmt"

type Quack struct{}

func (q *Quack) quack() {
	fmt.Println("Quack")
}

type Squeak struct{}

func (s *Squeak) quack() {
	fmt.Println("Squeak")
}

type MuteQuack struct{}

func (m *MuteQuack) quack() {
	fmt.Println("<< Silence >>")
}
