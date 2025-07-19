package main

type Waitress struct {
	allMenus MenuComponent
}

func NewWaitress(allMenus MenuComponent) *Waitress {
	return &Waitress{
		allMenus: allMenus,
	}
}

func (w *Waitress) PrintMenu() string {
	return w.allMenus.Print()
}

