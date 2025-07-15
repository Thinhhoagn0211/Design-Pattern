package main

type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) Duck {
	return &TurkeyAdapter{turkey: turkey}
}

func (a *TurkeyAdapter) Quack() {
	a.turkey.Gobble()
}

func (a *TurkeyAdapter) Fly() {
	// Turkeys fly short distances, so we can just call their Fly method
	for i := 0; i < 5; i++ {
		a.turkey.Fly()
	}
}
