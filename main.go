package main

import "fmt"

type Printer interface {
	Print(string)
}

type HP struct{}

func (d HP) Print(str string) {
	fmt.Println(str)
}

func start(q Printer) {
	q.Print("go interface")
}
func main() {
	myDuck := HP{}
	start(myDuck)
}
