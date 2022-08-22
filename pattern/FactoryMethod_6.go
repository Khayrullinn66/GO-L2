package main

import (
	"fmt"
	"log"
)

type Product interface {
	Name() string
}

// concrete product 1
type Pineapple struct {
}

func NewPinewapple() Product {
	return &Pineapple{}
}

func (a *Pineapple) Name() string {
	return "pineapple"
}

// concrete product 2
type Apple struct {
}

func NewApple() Product {
	return &Apple{}
}

func (a *Apple) Name() string {
	return "apple"
}

// concrete product 3
type Cherry struct {
}

func NewCherry() Product {
	return &Cherry{}
}

func (c *Cherry) Name() string {
	return "cherry"
}

type Creator interface {
	CreateProduct(str string) Product
}

type ConcreteCreator struct {
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (c *ConcreteCreator) CreateProduct(str string) Product {
	var product Product
	switch str {
	case "apple":
		product = &Apple{}
	case "pineapple":
		product = &Pineapple{}
	case "cherry":
		product = &Cherry{}
	default:
		log.Fatal("unknown product")
	}
	return product
}

func main() {
	c := NewCreator()
	fmt.Println(c.CreateProduct("apple").Name())
	fmt.Println(c.CreateProduct("cherry").Name())
	fmt.Println(c.CreateProduct("pineapple").Name())
	fmt.Println(c.CreateProduct("strawberry").Name())
}
