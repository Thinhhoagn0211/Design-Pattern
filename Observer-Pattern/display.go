package main

import "fmt"

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

func NewCurrentConditionsDisplay(weatherData Subject) *CurrentConditionsDisplay {
	display := &CurrentConditionsDisplay{
		weatherData: weatherData,
	}
	weatherData.registerObserver(display)
	return display
}

func (c *CurrentConditionsDisplay) update(temp float64, humidity float64, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.display()
}

func (c *CurrentConditionsDisplay) display() {
	fmt.Printf("Current conditions: %.1f°F degrees and %.1f%% humidity\n", c.temperature, c.humidity)
}

type StatisticsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

func NewStatisticsDisplay(weatherData Subject) *StatisticsDisplay {
	display := &StatisticsDisplay{
		weatherData: weatherData,
	}
	weatherData.registerObserver(display)
	return display
}

func (c *StatisticsDisplay) update(temp float64, humidity float64, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.display()
}

func (c *StatisticsDisplay) display() {
	fmt.Printf("Statistics conditions: %.1f°F degrees and %.1f%% humidity\n", c.temperature, c.humidity)
}

type ForecastDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

func NewForecastDisplay(weatherData Subject) *ForecastDisplay {
	display := &ForecastDisplay{
		weatherData: weatherData,
	}
	weatherData.registerObserver(display)
	return display
}

func (c *ForecastDisplay) update(temp float64, humidity float64, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.display()
}

func (c *ForecastDisplay) display() {
	fmt.Printf("Forecast conditions: %.1f°F degrees and %.1f%% humidity\n", c.temperature, c.humidity)
}
