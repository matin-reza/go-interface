package main

import "fmt"

type Printer interface {
	Print(string)
}

type Duck struct{}

func (d Duck) Print(str string) {
	fmt.Println(str)
}

func start(q Printer) {
	q.Print("go interface")
}
func main() {
	myDuck := Duck{}
	start(myDuck)
}
