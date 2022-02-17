package main

import "fmt"

// IProcess interface
type IProcess interface {
	process()
}

// Adapter struct
type Adapter struct {
	adaptee Adaptee
}

// Adapter class method process
func (a Adapter) process() {
	fmt.Println("Adapter process")
	a.adaptee.convert()
}

// Adaptee struct
type Adaptee struct {
	adapterType int
}

// Adaptee class method convert
func (a Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

// main method
func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
