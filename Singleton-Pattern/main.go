package main

import (
	"sync"
)

func main() {
	boiler := GetChocolateBoilerInstance()
	boiler.Fill()
	boiler.Boil()
	boiler.Drain()

	// Simulate concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instance := GetChocolateBoilerInstance()
			instance.Fill()
			instance.Boil()
			instance.Drain()
		}()
	}
	wg.Wait()	
}
