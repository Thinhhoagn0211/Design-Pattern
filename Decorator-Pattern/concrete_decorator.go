package main

type Mocha struct {
	beverage Beverage
}

func NewMocha(beverage Beverage) *Mocha {
	return &Mocha{
		beverage: beverage,
	}
}

func (m *Mocha) getDescription() string {
	return m.beverage.getDescription() + ", Mocha"
}

func (m *Mocha) getSize() string {
	return m.beverage.getSize()
}

func (m *Mocha) setSize(size string) {
	m.beverage.setSize(size)
}

func (m *Mocha) cost() float64 {
	if m.beverage.getSize() == "Large" {
		return m.beverage.cost() + 0.30
	} else if m.beverage.getSize() == "Medium" {
		return m.beverage.cost() + 0.25
	} else if m.beverage.getSize() == "Small" {
		return m.beverage.cost() + 0.15
	}
	return m.beverage.cost()
}

type Whip struct {
	beverage Beverage
}

func NewWhip(beverage Beverage) *Whip {
	return &Whip{
		beverage: beverage,
	}
}

func (w *Whip) getDescription() string {
	return w.beverage.getDescription() + ", Whip"
}

func (w *Whip) cost() float64 {
	if w.beverage.getSize() == "Large" {
		return w.beverage.cost() + 0.30
	} else if w.beverage.getSize() == "Medium" {
		return w.beverage.cost() + 0.25
	} else if w.beverage.getSize() == "Small" {
		return w.beverage.cost() + 0.15
	}
	return w.beverage.cost()
}

func (w *Whip) getSize() string {
	return w.beverage.getSize()
}

func (w *Whip) setSize(size string) {
	w.beverage.setSize(size)
}
