package main

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
}
