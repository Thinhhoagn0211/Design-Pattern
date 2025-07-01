package main

type Espresso struct {
	description string
	size        string
}

func NewEspresso() Beverage {
	return &Espresso{
		description: "Espresso",
	}
}

func (e *Espresso) getDescription() string {
	return e.description
}

func (e *Espresso) setSize(size string) {
	e.size = size
}

func (e *Espresso) getSize() string {
	return e.size
}

func (e *Espresso) cost() float64 {
	return 1.99
}

type HouseBlend struct {
	description string
	size        string
}

func NewHouseBlend() Beverage {
	return &HouseBlend{
		description: "House Blend Coffee",
	}
}

func (h *HouseBlend) getDescription() string {
	return h.description
}

func (h *HouseBlend) setSize(size string) {
	h.size = size
}

func (h *HouseBlend) getSize() string {
	return h.size
}

func (h *HouseBlend) cost() float64 {
	return 0.89
}

type DarkRoast struct {
	description string
	size        string
}

func NewDarkRoast() Beverage {
	return &DarkRoast{
		description: "Dark Roast Coffee",
	}
}

func (d *DarkRoast) getDescription() string {
	return d.description
}

func (d *DarkRoast) cost() float64 {
	return 0.99
}

func (d *DarkRoast) setSize(size string) {
	d.size = size
}

func (d *DarkRoast) getSize() string {
	return d.size
}
