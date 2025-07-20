package main

type GumballMachine struct {
	soldOutState State
	noQuarterState State
	hasQuarterState State
	soldState State
	winnerState State
	state State
	count int
}

func NewGumballMachine(count int) *GumballMachine {
	g := &GumballMachine{count: count}

	g.soldOutState = NewSoldOutState(g)
	g.noQuarterState = NewNoQuarterState(g)
	g.hasQuarterState = NewHasQuarterState(g)
	g.soldState = NewSoldState(g)
	g.winnerState = NewWinnerState(g)
	if count > 0 {
		g.state = g.noQuarterState
	} else {
		g.state = g.soldOutState
	} 

	return g
}

func (g *GumballMachine) InsertQuarter() {
	g.state.InsertQuarter()
}

func (g *GumballMachine) EjectQuarter() {
	g.state.EjectQuarter()
}

func (g *GumballMachine) TurnCrank() {
	g.state.TurnCrank()
	g.state.Dispense()
}

func (g *GumballMachine) setState(state State) {
	g.state = state
}

func (g *GumballMachine) getState() State {
	return g.state
}

func (g *GumballMachine) getNoQuarterState() State {
	return g.noQuarterState
}

func (g *GumballMachine) getHasQuarterState() State {
	return g.hasQuarterState
}

func (g *GumballMachine) getSoldState() State {
	return g.soldState
}

func (g *GumballMachine) getSoldOutState() State {
	return g.soldOutState
}

func (g *GumballMachine) getWinnerState() State {
	return g.winnerState
}

func (g *GumballMachine) getCount() int {
	return g.count
}

func (g *GumballMachine) getRandomNumber() int {
	// This function should return a random number, but for simplicity,
	// we will return a fixed value here.
	return 1 // Simulating a successful dispense
}

func (g *GumballMachine) releaseBall() {
	if g.count > 0 {
		g.count--
	}
}



