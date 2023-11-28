package main

import "fmt"

// DisplayElement 顯示元素
type DisplayElement interface {
	Display()
}

// Observer 觀察者
type Observer interface {
	Update(temp, humidity, pressure float32)
	DisplayElement
}

// Subject 主題
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

// WeatherData 氣象資料
type WeatherData struct {
	observers []Observer
	temp      float32
	humidity  float32
	pressure  float32
}

// RegisterObserver 註冊觀察者
func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

// RemoveObserver 移除觀察者
func (wd *WeatherData) RemoveObserver(o Observer) {
	for i, observer := range wd.observers {
		if observer == o {
			wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers 通知觀察者
func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observers {
		observer.Update(wd.temp, wd.humidity, wd.pressure)
	}
}

// MeasurementsChanged 氣象資料改變
func (wd *WeatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}

// SetMeasurements 設定氣象資料
func (wd *WeatherData) SetMeasurements(temp, humidity, pressure float32) {
	wd.temp = temp
	wd.humidity = humidity
	wd.pressure = pressure
	wd.MeasurementsChanged()
}

// CurrentConditionsDisplay 目前狀態顯示
type CurrentConditionsDisplay struct {
	temp     float32
	humidity float32
	weather  Subject
}

// NewCurrentConditionsDisplay 建立目前狀態顯示
func NewCurrentConditionsDisplay(weather Subject) *CurrentConditionsDisplay {
	ccd := &CurrentConditionsDisplay{
		weather: weather,
	}
	weather.RegisterObserver(ccd)
	return ccd
}

// Update 更新資料
func (ccd *CurrentConditionsDisplay) Update(temp, humidity, pressure float32) {
	ccd.temp = temp
	ccd.humidity = humidity
	ccd.Display()
}

// Display 顯示資料
func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Println("temp:", ccd.temp)
	fmt.Println("humidity:", ccd.humidity)
}

func main() {
	weatherData := &WeatherData{}
	currentConditionsDisplay := NewCurrentConditionsDisplay(weatherData)
	currentConditionsDisplay.Display()
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}
