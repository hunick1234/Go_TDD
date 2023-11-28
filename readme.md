練習設計模式並使用 TDD 進行開發

## 練習題目

- 貪吃蛇遊戲

## 設計模式

> 封裝:簡單理解-資料隱藏\
> 物件內部資料成員的存取限制\
> 但封裝不只是資料隱藏，可泛指各種隱藏
> _p.17_

## 1. 策略模式

### 設計一隻鴨子

- 從鴨子向上看

  ```go
  type Duck struct {

  }

  func (d *Duck) Swim() {
    fmt.Println("I can swim")
  }

  func (d *Duck) Fly() {
    fmt.Println("I can fly")
  }

  func (d *Duck) Quack() {
    fmt.Println("I can quack")
  }

  func (d *Duck) Display() {
    fmt.Println("I am a duck")
  }

  type MyDuck struct {
    Duck
  }

  // if this duck can't fly
  type CantFlyDuck struct {
    Duck
  }

  func (d *CantFlyDuck) Fly() {
    fmt.Println("I can't fly")
  }
  // if this duck can't quack
  type CantQuackDuck struct {
    Duck
  }


  func (d *CantQuackDuck) Fly() {
    fmt.Println("I can't quack")
  }

  ```

在上敘程式中
他有基本的鴨子行為 但是如果我想要新增一隻不會飛的鴨子\
我就得為他新增一個 struct 並覆寫 Fly()\
如果我想要新增一隻不會叫的鴨子\
我就得為他新增一個 struct 並覆寫 Quack()\
這會造成程式碼的膨脹\
且如果我想要新增一隻不會飛也不會叫的鴨子\
我就得為他新增一個 struct 並覆寫 Fly() 和 Quack()

我們變更行為時
我們會得到很多 重複的**_行為_**散佈在各個 struct method 中\
得直接對類別進行修改\
這違反了開放封閉原則（Open-Closed Principle）\
對擴展開放，對修改封閉

- 從一隻鴨子向下看

  ```go
  package main

  import "fmt"

  type Flyer interface {
      Fly()
  }

  type Quacker interface {
      Quack()
  }

  type Swimer interface {
      Swim()
  }

  type Duck struct {
      flyBehavior   Flyer
      quackBehavior Quacker
      swimBehavior  Swimer
  }

  func (d *Duck) performFly() {
      d.flyBehavior.Fly()
  }

  func (d *Duck) performQuack() {
      d.quackBehavior.Quack()
  }

  func (d *Duck) performSwim() {
      d.swimBehavior.Swim()
  }

  // CantFly 實現了 Flyer 但是代表不會飛的行為
  type CantFly struct{}

  func (cf *CantFly) Fly() {
      fmt.Println("I can't fly")
  }

  // CanQuack 實現了 Quacker 接口
  type CanQuack struct{}

  func (cq *CanQuack) Quack() {
      fmt.Println("Quack")
  }

  // CanSwim 實現了 Swimer 接口
  type CanSwim struct{}

  func (cs *CanSwim) Swim() {
      fmt.Println("Swimming")
  }

  // MyDuck 是一個自定義的鴨子類型
  type MyDuck struct {
      Duck
  }

  func main() {
      myDuck := MyDuck{
          Duck{
              flyBehavior:   new(CantFly),
              quackBehavior: new(CanQuack),
              swimBehavior:  new(CanSwim),
          },
      }

      myDuck.performFly()   // 輸出：I can't fly
      myDuck.performQuack() // 輸出：Quack
      myDuck.performSwim()  // 輸出：Swimming
  }

  ```

  我們 duck struct 從 IS-A 鴨子(是一個) 變成 HAS-A 鴨子(有一個)\
  他是有鴨子行為的任何東西

  ```go
  // 我們還可以動態的改變行為
  func (d *Duck) SetFlyBehavior(fb Flyer) {
      d.flyBehavior = fb
  }
  myDuck.SetFlyBehavior(new(CanFly))

  ```

##### WHY? 為什麼要這樣設計?

當這樣設計時 我鴨子每有一種行為變化 我就得為她新增一個 struct or class 去實作其介面 程式碼不就會變出很多 不必要的類別嗎?\

##### So Why?!

上敘例子 其實就是使用策略模式
的確，有機會導致許多小的類別或結構體被創建來實現不同的行為\

- 優點

  - 可擴展性：當新的行為需要被添加到系統中時，不需要修改現有的類別。這遵循了開放封閉原則（Open-Closed Principle），系統對擴展開放，對修改封閉。
  - 可維護性：每個行為都被封裝在自己的類別中，使得代碼易於理解和維護。
  - 解耦合：不同的行為和對象之間解耦，提高了代碼的靈活性。

- 缺點
  - 類別數量增加：每個策略都是一個類別，如果策略過多，類別數量將會增加。
  - 複雜性增加：雖然每個類別都很簡單，但是大量的類別可能使得系統整體複雜性增加。

#### 解決方案

- 評估需求：在設計時，先評估是否真的需要這麼多的行為變化。如果系統不太可能經常變更或擴展，那麼一個更簡單的設計可能更合適。

- 重用和組合：尋找重用現有行為的機會。有時，通過組合已有的行為，可以創造出新的行為，而不是創建一個全新的類別。

- 函數類型：在某些情況下，特別是行為只涉及單一方法時，可以考慮使用 Go 的函數類型而不是接口。這樣可以直接賦予函數到結構體，而不需要為每種行為創建一個新的類別。

- 接口內嵌：如果多個行為經常一起出現，可以考慮將接口內嵌在一起，這樣可以通過實現一個較大的接口來簡化某些情況。

> **使用策略模式（Strategy Pattern）來管理行為(演算法)**

---


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
  // DisplayElement 顯示元素
  type DisplayElement interface {
      Display()
  }

  // Observer 觀察者
  type Observer interface {
      Update(temp, humidity, pressure float32)
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
      humidity float
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
      var weatherData Subject
      weatherData := &WeatherData{}
      currentConditionsDisplay := NewCurrentConditionsDisplay(weatherData)
      weatherData.SetMeasurements(80, 65, 30.4)
      weatherData.SetMeasurements(82, 70, 29.2)
      weatherData.SetMeasurements(78, 90, 29.2)
  }

  ```
  ```

##### So Why?!


- 優點
  - 解耦合：主題和觀察者之間解耦，使得主題和觀察者可以獨立地改變和重用。
  - 可維護性：主題或觀察者可以獨立地擴展，而不會影響到另一方。
  - 可重用性：主題和觀察者可以獨立地重用。
  - 低耦合：主題和觀察者之間的依賴關係是抽象的，而不是具體的類別。

- 缺點
  


## 3. 裝飾者模式

```go
package main

import "fmt"

type Beverager interface {
	Cost() float32
	GetDescription() string
}

type HouseBlend struct {
	description string
}

func (hb *HouseBlend) Cost() float32 {
	return 0.89
}

func (hb *HouseBlend) GetDescription() string {
	return hb.description
}

type DarkRoast struct {
	description string
}

func (dr *DarkRoast) Cost() float32 {
	return 0.99
}

func (dr *DarkRoast) GetDescription() string {
	return dr.description
}

type Espresso struct {
	description string
}

func (e *Espresso) Cost() float32 {
	return 1.99
}

func (e *Espresso) GetDescription() string {
	return e.description
}

type Mocha struct {
	beverage Beverager
}

func (m *Mocha) Cost() float32 {
	return 0.2 + m.beverage.Cost()
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func NewMocha(beverage Beverager) *Mocha {
	return &Mocha{
		beverage: beverage,
	}
}

type Milk struct {
	beverage Beverager
}

func (m *Milk) Cost() float32 {
	return 0.3 + m.beverage.Cost()
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", Milk"
}

func NewMilk(beverage Beverager) *Milk {
	return &Milk{
		beverage: beverage,
	}
}

func main() {
	beverage := &Espresso{description: "Espresso"}
	fmt.Println(beverage.GetDescription(), beverage.Cost())

	var beverage2 Beverager
	beverage2 = &DarkRoast{description: "DarkRoast"}
	beverage2 = NewMocha(beverage2)
	beverage2 = NewMocha(beverage2)
	beverage2 = NewMilk(beverage2)
	fmt.Println(beverage2.GetDescription(), beverage2.Cost())
}

```



## How to TDD

### TDD Cycle

1.  Add a test
2.  Run all tests and see if the new one fails
3.  Write some code
4.  Run tests
5.  Refactor code
6.  Repeat

## PKG Testing

- t testing.T
  - t.Run()
  - t.Error()
  - t.Errorf()
