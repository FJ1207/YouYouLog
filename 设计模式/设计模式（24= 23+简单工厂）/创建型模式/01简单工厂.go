package main

import "fmt"

// ------ 抽象层--------
type Fruit interface {
	show()
}

// ------- 实现层-------
type Banana struct {
}

func (b Banana) show() {
	fmt.Println("我是 banana")
}

type Apple struct {
}

func (a Apple) show() {
	fmt.Println("我是 apple")
}

type Pear struct{}

func (p Pear) show() {
	fmt.Println("我是 pear")
}

// -------简单工厂--------
type Factory struct {
}

func (f Factory) CreateFruitFactory(name string) Fruit { // 返回抽象 不符合开闭原则
	var fruit Fruit
	if name == "pear" {
		fruit = Pear{}
	} else if name == "apple" {
		fruit = Apple{}
	} else if name == "banana" {
		fruit = Banana{}
	}
	return fruit
}

// ------- 业务逻辑层------
func main() {}
