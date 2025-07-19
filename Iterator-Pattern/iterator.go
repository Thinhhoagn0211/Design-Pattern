package main


type DinerMenuIterator struct {
	items [6]*MenuItem
	position int
}

func NewDinerMenuIterator(items [6]*MenuItem) *DinerMenuIterator {
	return &DinerMenuIterator{
		items: items,
		position: 0,
	}
}

func (d *DinerMenuIterator) HasNext() bool {
	return d.position < len(d.items)
}

func (d *DinerMenuIterator) Next() *MenuItem {
	if !d.HasNext() {
		return &MenuItem{} // Return an empty MenuItem if no more items are available
	}
	item := d.items[d.position]
	d.position++
	return item
}

type PancakeHouseMenuIterator struct {
	items []*MenuItem
	position int
}

func NewPancakeHouseMenuIterator(items []*MenuItem) *PancakeHouseMenuIterator {
	return &PancakeHouseMenuIterator{
		items: items,
		position: 0,
	}
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	return p.position < len(p.items)
}

func (p *PancakeHouseMenuIterator) Next() *MenuItem {
	if !p.HasNext() {
		return &MenuItem{} // Return an empty MenuItem if no more items are available
	}
	item := p.items[p.position]
	p.position++
	return item
}

