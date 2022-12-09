package main

// ------抽象层-------
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

// ------基础模块层-------

func main() {

}
