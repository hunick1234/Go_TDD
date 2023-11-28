
# 2. 觀察者模式

- 設計一個氣象站

  - 顯示目前狀態
  - 顯示氣象統計資料
  - 顯示氣象預測資料

  ```go
  type WeatherData struct {
    observers []CurrentConditionsDisplay
    temp      float32
    humidity  float32
    pressure  float32
  }

    // CurrentConditionsDisplay 目前狀態顯示
  type CurrentConditionsDisplay struct {
      temp     float32
      humidity float
      weather  Subject
  }

  // NewCurrentConditionsDisplay 建立目前狀態顯示
  func NewCurrentConditionsDisplay(weather wheatherData) *CurrentConditionsDisplay {
      ccd := &CurrentConditionsDisplay{
          weather: weather,
      }
      weather.observers = append(weather.observers, ccd)
      return ccd
  }

  // Update 更新資料
  func (ccd *CurrentConditionsDisplay) Update(temp, humidity, pressure float32) {
      ccd.temp = temp
      ccd.humidity = humidity
      ccd.Display()
  }

  func (ccd *CurrentConditionsDisplay)Display() {
    fmt.Println("temp:", ccd.temp)
    fmt.Println("humidity:", ccd.humidity)
  }

  ```

  上面也實現了觀察者模式的基本概念
  但兩者間的耦合性太高
  當我們想要修改時 一切都很麻煩

  所以我們用介面來解耦合
  subject & observer interface

```go
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
	//var weatherData Subject
	weatherData := &WeatherData{}
	currentConditionsDisplay := NewCurrentConditionsDisplay(weatherData)
	currentConditionsDisplay.Display()
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)

}


```



##### So Why?!

- 優點

  - 解耦合：主題和觀察者之間解耦，使得主題和觀察者可以獨立地改變和重用。
  - 可維護性：主題或觀察者可以獨立地擴展，而不會影響到另一方。
  - 可重用性：主題和觀察者可以獨立地重用。
  - 低耦合：主題和觀察者之間的依賴關係是抽象的，而不是具體的類別。

- 缺點
