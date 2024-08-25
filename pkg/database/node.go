package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type Node struct {
	ID       int
	Name     string
	ParentID *int
	Path     string
	Children []*Node
}

func CreateNode(db *sql.DB, name string) {
	_, err := db.Exec("INSERT INTO nodes (name, parent_id, full_path) VALUES (?, ?, ?);", name, nil, "/")
	if err != nil {
		fmt.Println(err)
	}
}

func AddChildNode(db *sql.DB, parentId int, name, path string) {
	_, err := db.Exec("INSERT INTO nodes (name, parent_id, full_path) VALUES (?, ?, ?);", name, parentId, path)
	if err != nil {
		fmt.Println(err)
	}
}

func GetChildren(db *sql.DB, fullPath string) []Node {
	rows, err := db.Query("SELECT id, name, parent_id, full_path FROM nodes WHERE full_path = ?", fullPath)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var nodes []Node
	for rows.Next() {
		var node Node
		if err := rows.Scan(&node.ID, &node.Name, &node.ParentID, &node.Path); err != nil {
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func GetNodeIDByName(db *sql.DB, name string) []Node {
	rows, err := db.Query("SELECT id FROM nodes WHERE name = ?", name)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var nodes []Node
	for rows.Next() {
		var node Node
		if err := rows.Scan(&node.ID); err != nil {
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func GetNodeIDByPath(db *sql.DB, name, path string) int {
	fmt.Printf("Name: %s, Path: %s\n", name, path)
	row, err := db.Query("SELECT id FROM nodes WHERE name = ? AND full_path = ?", name, path)
	if err != nil {
		fmt.Println(err)
	}
	var node Node
	for row.Next() {
		row.Scan(&node.ID)
		fmt.Println(node)
	}
	return node.ID
}

func CompleteNode() {
}

func GetNode(db *sql.DB, id int) Node {
	row := db.QueryRow("SELECT id, name, parent_id, full_path FROM nodes WHERE id=?", id)
	node := Node{}
	row.Scan(&node.ID, &node.Name, &node.ParentID, &node.Path)
	return node
}

func GetNodes(db *sql.DB) []Node {
	rows, _ := db.Query("SELECT id, name, parent_id from nodes WHERE parent_id IS NULL")
	defer rows.Close()
	var nodes []Node
	for rows.Next() {
		var node Node
		if err := rows.Scan(&node.ID, &node.Name, &node.ParentID); err != nil {
		}
		nodes = append(nodes, node)
	}

	return nodes

}

func NodeWalk(db *sql.DB, input string) (int, error) {
	nodesParsed := strings.Split(input, "/")
	//for _, node := range nodesParsed {
	//	fmt.Println(node)
	//}
	ids := GetNodeIDByName(db, nodesParsed[len(nodesParsed)-1])
	var nodeID int
	var err error
	switch {
	case len(ids) == 0:
		err = fmt.Errorf("node: %s was not found", nodesParsed[0])
	case len(ids) == 1:
		nodeID = ids[0].ID
	case len(ids) > 1:
		//err = fmt.Errorf("Multiple nodes with name: %s found. Please use a path", nodesParsed[0])
		nodeID = ids[0].ID
	}
	return nodeID, err
}

func UpdateNode() {
}

func DeleteNode(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM nodes where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
}
