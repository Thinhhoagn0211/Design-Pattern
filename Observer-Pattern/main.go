package main

import "time"

func main() {
	weatherData := NewWeatherData()

	// currentDisplay := NewCurrentConditionsDisplay(weatherData)
	// statisticsDisplay := NewStatisticsDisplay(weatherData)
	// forecastDisplay := NewForecastDisplay(weatherData)
	currentDisplay := NewCurrentConditionsDisplay(weatherData)
	statisticsDisplay := NewStatisticsDisplay(weatherData)
	forecastDisplay := NewForecastDisplay(weatherData)

	weatherData.setMeasurements(80, 65, 20.4)
	weatherData.setMeasurements(82, 70, 29.2)
	weatherData.setMeasurements(78, 90, 29.2)

	time.Sleep(2 * time.Second)
	weatherData.removeObserver(currentDisplay)
	weatherData.setMeasurements(70, 50, 29.2)
	time.Sleep(2 * time.Second)
	weatherData.removeObserver(statisticsDisplay)
	weatherData.setMeasurements(60, 40, 29.2)
	time.Sleep(2 * time.Second)
	weatherData.removeObserver(forecastDisplay)
	weatherData.setMeasurements(50, 30, 29.2)
	time.Sleep(2 * time.Second)
}
