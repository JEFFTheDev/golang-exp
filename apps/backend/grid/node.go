package grid

import "fmt"

func DoSomething() {
	fmt.Println("AAA")
}

type NodeBehaviour interface {
}

type Node struct {
	Id     int
	X      int
	Y      int
	gridId int
}
