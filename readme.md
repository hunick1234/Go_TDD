練習設計模式並使用 TDD 進行開發

## 練習題目

- 貪吃蛇遊戲

## 設計模式

> 封裝:簡單理解-資料隱藏\
> 物件內部資料成員的存取限制\
> 但封裝不只是資料隱藏，可泛指各種隱藏
> _p.17_

## 1. 策略模式
[strategy](./About_Design_Patten/Strategy.md)

## 2. 觀察者模式
[observer](./About_Design_Patten/Observer.md)

## 3. 裝飾者模式
[decorator](./About_Design_Patten/Decorator.md)

## 4. 工廠模式
[factory](./About_Design_Patten/Factory.md)

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
