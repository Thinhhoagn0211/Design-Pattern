package main


type BasePizzaStore struct {
	Factory func(string) Pizza
}

func (store *BasePizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := store.Factory(pizzaType)
	if pizza == nil {
		return nil
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

type  NYPizzaStore struct {
	BasePizzaStore
}

func NewNYPizzaStore() PizzaStore {
	return &NYPizzaStore{
		BasePizzaStore{
			Factory: func(pizzaType string) Pizza {
				switch pizzaType {
				case "cheese":
					return NewNYStyleCheesePizza()
				case "pepperoni":
					return NewNYStylePepperoniPizza()
				case "veggie":
					return NewNYStyleVeggiePizza()
				default:
					return nil
				}
			},
		},
	}
}

type ChicagoPizzaStore struct {
	BasePizzaStore
}

func NewChicagoPizzaStore() PizzaStore {	
	return &ChicagoPizzaStore{
		BasePizzaStore{
			Factory: func(pizzaType string) Pizza {
				switch pizzaType {
				case "cheese":
					return NewChicagoStyleCheesePizza()
				case "pepperoni":
					return NewChicagoStylePepperoniPizza()
				case "veggie":
					return NewChicagoStyleVeggiePizza()
				default:
					return nil
				}
			},
		},
	}
}

