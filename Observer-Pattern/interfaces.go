package main

type Subject interface {
	registerObserver(o Observer)
	removeObserver(o Observer)
	notifyObservers()
}

type Observer interface {
	update(temp float64, humidity float64, pressure float64)
}

type DisplayElement interface {
	display()
}
