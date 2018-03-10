package main

import (
	"fmt"
)

type Door interface {
	getDescription()
}

type FittingExpert interface {
	getDescription()
}

type DoorFactory interface {
	makeDoor() Door
	makeFittingExpert() FittingExpert
}

type WoodenDoorFactory struct{}

func (wdf *WoodenDoorFactory) makeDoor() Door {
	return &WoodenDoor{}
}

func (wdf *WoodenDoorFactory) makeFittingExpert() FittingExpert {
	return &Carpenter{}
}

type WoodenDoor struct{}

func (wd *WoodenDoor) getDescription() {
	fmt.Println("I'm a wooden door.")
}

type IronDoorFactory struct{}

func (idf *IronDoorFactory) makeDoor() Door {
	return &IronDoor{}
}

func (id *IronDoorFactory) makeFittingExpert() FittingExpert {
	return &Welder{}
}

type IronDoor struct{}

func (id *IronDoor) getDescription() {
	fmt.Println("I'm an iron door.")
}

type Carpenter struct{}

func (c *Carpenter) getDescription() {
	fmt.Println("I can only fix wooden door.")
}

type Welder struct{}

func (w *Welder) getDescription() {
	fmt.Println("I can only fix iron doors.")
}

func main() {
	woodenDoorFactory := WoodenDoorFactory{}
	ironDoorFactory := IronDoorFactory{}

	woodenDoor := woodenDoorFactory.makeDoor()
	ironDoor := ironDoorFactory.makeDoor()

	woodenDoor.getDescription() // I'm a wooden door.
	ironDoor.getDescription()   // I'm an iron door.

	woodenDoorExpert := woodenDoorFactory.makeFittingExpert()
	ironDoorExpert := ironDoorFactory.makeFittingExpert()

	woodenDoorExpert.getDescription() // I can only fix wooden door.
	ironDoorExpert.getDescription()   // I can only fix iron doors.
}
