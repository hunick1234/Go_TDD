
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
