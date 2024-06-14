package main

import "fmt"

type abstractFactory interface {
	make() IObject
}

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

type AFactory struct{}

func (f AFactory) make() IObject {
	return objectA{}
}

type BFactory struct{}

func (f BFactory) make() IObject {
	return objectB{}
}

func main() {
	typesOfThing := []string{"a", "b", "c"}
	for _, t := range typesOfThing {
		f, err := getFactoryFor(t)
		if err != nil {
			fmt.Printf("making a %s: %v", t, err)
			continue
		}
		o := f.make()
		o.doThing()
	}
}

func getFactoryFor(thing string) (abstractFactory, error) {
	switch thing {
	case "a":
		return AFactory{}, nil
	case "b":
		return BFactory{}, nil
	default:
		return nil, fmt.Errorf("I cannot make a factory to make a %s", thing)
	}
}
