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
我們會得到很多 重複的***行為***散佈在各個 struct method 中\
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
  
  我們duck struct 從IS-A 鴨子(是一個) 變成 HAS-A 鴨子(有一個)\
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

  * 優點
    * 可擴展性：當新的行為需要被添加到系統中時，不需要修改現有的類別。這遵循了開放封閉原則（Open-Closed Principle），系統對擴展開放，對修改封閉。
    * 可維護性：每個行為都被封裝在自己的類別中，使得代碼易於理解和維護。
    * 解耦合：不同的行為和對象之間解耦，提高了代碼的靈活性。

  * 缺點
    * 類別數量增加：每個策略都是一個類別，如果策略過多，類別數量將會增加。
    * 複雜性增加：雖然每個類別都很簡單，但是大量的類別可能使得系統整體複雜性增加。


#### 解決方案
  - 評估需求：在設計時，先評估是否真的需要這麼多的行為變化。如果系統不太可能經常變更或擴展，那麼一個更簡單的設計可能更合適。

  - 重用和組合：尋找重用現有行為的機會。有時，通過組合已有的行為，可以創造出新的行為，而不是創建一個全新的類別。

  - 函數類型：在某些情況下，特別是行為只涉及單一方法時，可以考慮使用 Go 的函數類型而不是接口。這樣可以直接賦予函數到結構體，而不需要為每種行為創建一個新的類別。

  - 接口內嵌：如果多個行為經常一起出現，可以考慮將接口內嵌在一起，這樣可以通過實現一個較大的接口來簡化某些情況。

> **使用策略模式（Strategy Pattern）來管理行為**

----

## How to TDD

### TDD Cycle

1.  Add a test
2.  Run all tests and see if the new one fails
3.  Write some code
4.  Run tests
5.  Refactor code
6.  Repeat

## PKG Testing

- t \*testing.T
  - t.Run()
  - t.Error()
  - t.Errorf()
