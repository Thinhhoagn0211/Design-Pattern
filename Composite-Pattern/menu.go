package main

type Menu struct {
	components []MenuComponent
	name		string
	description string
}

func NewMenu(name, description string) MenuComponent {
	return &Menu{
		name:        name,
		description: description,
	}
}

func (m *Menu) Add(component MenuComponent) {
	m.components = append(m.components, component)
}

func (m *Menu) Remove(component MenuComponent) {
	for i, c := range m.components {
		if c == component {
			m.components = append(m.components[:i], m.components[i+1:]...)
			break
		}
	}
}

func (m *Menu) GetChild(index int) MenuComponent {
 	if index < 0 || index >= len(m.components) {
		return nil
	}
	return m.components[index]
}

func (m *Menu) GetName() string {
	return m.name
}

func (m *Menu) GetDescription() string {
	return m.description
}

func (m *Menu) Print() string {
	result := m.name + ": " + m.description + "\n"
	for _, component := range m.components {
		result += "  " + component.Print() + "\n"
	}
	return result
}

func (m *Menu) GetPrice() float64 {
	// Menus do not have a price, so we return 0.0
	return 0.0
}

func (m *Menu) IsVegetarian() bool {
	// Menus do not have a vegetarian status, so we return false
	return false
}



