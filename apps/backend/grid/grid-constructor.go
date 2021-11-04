package grid

import "fmt"

func GetGridWithNodesAndEdgesById(id int) *Grid {

	// First get the grid record from the database.
	var gridDb = GetGridById(id)
	if gridDb == nil {
		fmt.Printf(fmt.Sprintf("No grid with Id %d found!", id))
		return nil
	}

	// Now get the nodes from the database.
	nodes := GetNodesOfGrid(gridDb.Id)

	// Now get the edges from the databases.
	edges := GetEdgesOfGrid(gridDb.Id)

	// Fill all edges with the correct nodes.
	for i := 0; i < len(*edges); i++ {
		edge := &(*edges)[i]
		for j := 0; j < len(*nodes); j++ {
			node := &(*nodes)[j]

			// The database call has mapped the Ids of the destination and source nodes, but not the rest of the properties.
			if edge.destination_id == node.Id {
				edge.Destination = *node
			}

			if edge.source_id == node.Id {
				edge.Source = *node
			}
		}
	}

	gridDb.Nodes = nodes
	gridDb.Edges = edges

	// Return the result now that all objects are filled.
	return gridDb
}
