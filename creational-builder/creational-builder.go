package main

import (
	"fmt"
)

type BurgerBuilder struct {
	size    int
	tomato  bool
	cheese  bool
	lettuce bool
}

func (bb *BurgerBuilder) addTomato() *BurgerBuilder {
	bb.tomato = true
	return bb
}

func (bb *BurgerBuilder) addCheese() *BurgerBuilder {
	bb.cheese = true
	return bb
}

func (bb *BurgerBuilder) addLettuce() *BurgerBuilder {
	bb.lettuce = true
	return bb
}

func (bb *BurgerBuilder) build() *Burger {
	burger := &Burger{}
	return burger.make(bb)
}

type Burger struct {
	size    int
	tomato  bool
	cheese  bool
	lettuce bool
}

func (b *Burger) make(bb *BurgerBuilder) *Burger {
	b.size = bb.size
	b.tomato = bb.tomato
	b.cheese = bb.cheese
	b.lettuce = bb.lettuce
	return b
}

func (b *Burger) String() string {
	return fmt.Sprintf("Burger (size: %v, tomato: %v, cheese: %v, lettuce: %v)", b.size, b.tomato, b.cheese, b.lettuce)
}

func main() {
	burger := (&BurgerBuilder{size: 15}).addTomato().addCheese().addLettuce().build()
	fmt.Println(burger)
}
