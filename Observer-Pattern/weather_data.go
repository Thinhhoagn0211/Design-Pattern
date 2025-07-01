package main

type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make([]Observer, 0),
	}
}

func (wd *WeatherData) registerObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) removeObserver(o Observer) {
	for i, observer := range wd.observers {
		if observer == o {
			wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
			break
		}
	}
}

func (wd *WeatherData) notifyObservers() {
	for _, observer := range wd.observers {
		observer.update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) measurementsChanged() {
	wd.notifyObservers()
}

func (wd *WeatherData) setMeasurements(temperature float64, humidity float64, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}
