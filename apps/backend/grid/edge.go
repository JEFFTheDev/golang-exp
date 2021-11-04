package grid

type EdgeBehaviour interface {
}

type Edge struct {
	Id             int
	Source         Node
	Destination    Node
	Cost           int
	destination_id int
	source_id      int
	gridId         int
}
