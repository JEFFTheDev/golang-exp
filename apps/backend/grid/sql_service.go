package grid

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetGridById(id int) *Grid {
	var sqlResult = executeQuery("SELECT * FROM Grid WHERE Id=?", id)

	var gridDb Grid

	for sqlResult.Next() {
		var err = sqlResult.Scan(&gridDb.Id, &gridDb.Name)
		if err != nil {
			fmt.Println("Error occurred: " + err.Error())
			return nil
		}

		fmt.Println(gridDb.Id)
		fmt.Println(gridDb.Name)
	}

	sqlResult.Close()

	return &gridDb
}

func GetEdgesOfGrid(gridId int) *[]Edge {
	var sqlResult = executeQuery("SELECT * FROM Edges WHERE grid_id=?", gridId)

	var edges []Edge

	for sqlResult.Next() {
		var edge Edge
		var err = sqlResult.Scan(&edge.destination_id, &edge.source_id, &edge.gridId, &edge.Id, &edge.Cost)
		if err != nil {
			fmt.Println("Error occurred: " + err.Error())
			return nil
		}

		edges = append(edges, edge)
	}

	sqlResult.Close()

	return &edges
}

func GetNodesOfGrid(gridId int) *[]Node {
	var sqlResult = executeQuery("SELECT * FROM Nodes WHERE grid_id=?", gridId)

	var nodes []Node

	for sqlResult.Next() {
		var node Node
		var err = sqlResult.Scan(&node.Id, &node.X, &node.Y, &node.gridId)
		if err != nil {
			fmt.Println("Error occurred: " + err.Error())
			return nil
		}

		nodes = append(nodes, node)
	}

	sqlResult.Close()

	return &nodes
}

func executeQuery(query string, args int) *sql.Rows {
	db, err := sql.Open("sqlite3", "./infrastructure.db")
	if err != nil {
		fmt.Println("Error occurred: " + err.Error())
		return nil
	}

	rows, err := db.Query(query, args)
	if err != nil {
		fmt.Println("Error occurred: " + err.Error())
		return nil
	}

	db.Close()

	return rows
}
