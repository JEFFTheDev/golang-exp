package grid

type GridBehaviour interface {
}

type Grid struct {
	Id    int
	Name  string
	Nodes *[]Node
	Edges *[]Edge
}
