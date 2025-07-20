package main

import (
	"fmt"
	"math/rand"
)

type NoQuarterState struct {
	gumballMachine *GumballMachine
}

func NewNoQuarterState(gumballMachine *GumballMachine) State {
	return &NoQuarterState{
		gumballMachine: gumballMachine,
	}
}

func (n *NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	n.gumballMachine.setState(n.gumballMachine.getHasQuarterState())
}

func (n *NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (n *NoQuarterState) TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

func (n *NoQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

type HasQuarterState struct {
	randomWinner int
	gumballMachine *GumballMachine
}

func NewHasQuarterState(gumballMachine *GumballMachine) State {
	return &HasQuarterState{
		gumballMachine: gumballMachine,
	}
}

func (h *HasQuarterState) InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

func (h *HasQuarterState) EjectQuarter() {
	fmt.Println("Quarter returned")
	h.gumballMachine.setState(h.gumballMachine.getNoQuarterState())
}

func (h *HasQuarterState) TurnCrank() {
	fmt.Println("You turned the crank")
	winner := rand.Intn(10) == 0
	if winner && h.gumballMachine.getCount() > 1 {
		fmt.Println("You are a winner! You get two gumballs for your quarter")
		h.gumballMachine.setState(h.gumballMachine.getWinnerState())
	} else {
		fmt.Println("No winner this time")
		h.gumballMachine.setState(h.gumballMachine.getSoldState())
	}
}

func (h *HasQuarterState) Dispense() {
	fmt.Println("No gumball dispensed")
}

type SoldState struct {
	gumballMachine *GumballMachine
}

func NewSoldState(gumballMachine *GumballMachine) State {
	return &SoldState{
		gumballMachine: gumballMachine,
	}
}

func (s *SoldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (s *SoldState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (s *SoldState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

func (s *SoldState) Dispense() {
	s.gumballMachine.releaseBall()	
	if s.gumballMachine.getCount() > 0 {
		s.gumballMachine.setState(s.gumballMachine.getNoQuarterState())
	} else {
		fmt.Println("Oops, out of gumballs!")
		s.gumballMachine.setState(s.gumballMachine.getSoldOutState())
	}
}

type SoldOutState struct {
	gumballMachine *GumballMachine
}

func NewSoldOutState(gumballMachine *GumballMachine) State {
	return &SoldOutState{
		gumballMachine: gumballMachine,
	}
}

func (s *SoldOutState) InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

func (s *SoldOutState) EjectQuarter() {
	fmt.Println("You can't eject, you haven't inserted a quarter")
}

func (s *SoldOutState) TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}

func (s *SoldOutState) Dispense() {
	fmt.Println("No gumball dispensed")
}

type WinnerState struct {
	gumballMachine *GumballMachine
}

func NewWinnerState(gumballMachine *GumballMachine) State {
	return &WinnerState{
		gumballMachine: gumballMachine,
	}
}

func (w *WinnerState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (w *WinnerState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (w *WinnerState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

func (w *WinnerState) Dispense() {
	w.gumballMachine.releaseBall()	
	if w.gumballMachine.getCount() == 0 {
		w.gumballMachine.setState(w.gumballMachine.getSoldOutState())
	} else {
		fmt.Println("You are a winner! You get an extra gumball")
		w.gumballMachine.releaseBall()
		if w.gumballMachine.getCount() > 0 {
			w.gumballMachine.setState(w.gumballMachine.getNoQuarterState())
		} else {
			w.gumballMachine.setState(w.gumballMachine.getSoldOutState())
		}
	}
}
