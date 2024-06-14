package main

import "fmt"

type factory struct{}

type IObject interface {
	doThing()
}

type objectA struct{}
type objectB struct{}

func (o objectA) doThing() {
	fmt.Println("I'm an A")
}

func (o objectB) doThing() {
	fmt.Println("I'm a B")
}

func (f factory) make(aThing string) (IObject, error) {
	switch aThing {
	case "a":
		return objectA{}, nil
	case "b":
		return objectB{}, nil
	default:
		return nil, fmt.Errorf("I don't know how to make a %s", aThing)
	}
}

func main() {
	factory := factory{}
	anA, _ := factory.make("a")
	aB, _ := factory.make("b")

	anA.doThing()
	aB.doThing()

	_, err := factory.make("c")
	if err != nil {
		fmt.Printf("creating a thing: %v", err)
	}
}
