package database

import (
	"database/sql"
	"fmt"
)

type Tag struct {
	Key   string
	Value string
}

func Taggy() {
	fmt.Println("TEST")
}

func AddTag(db *sql.DB, nodeID, tagID int) {
	db.Exec("INSERT INTO node_tags (node_id, tag_id) VALUES (?, ?)", nodeID, tagID)
}

func RemoveTag(db *sql.DB, nodeID, tagID int) {
	db.Exec("DELETE FROM node_tags WHERE node_id = ? AND tag_id = ?;", nodeID, tagID)
}

func CreateTag(db *sql.DB, key, value string) {
	db.Exec("INSERT INTO tags (key, value) VALUES (?, ?)", key, value)
}

func GetTags(db *sql.DB) {
	rows, _ := db.Query("SELECT key, value FROM tags")

	defer rows.Close()
	var tags []Tag
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.Key, &tag.Value); err != nil {
		}
		tags = append(tags, tag)
	}
	for _, tag := range tags {
		fmt.Println(tag)
	}
}

func GetTagIDByName(db *sql.DB, tagKey, tagValue string) int {
	row := db.QueryRow("SELECT id FROM tags WHERE key = ? AND value = ?", tagKey, tagValue)
	var tag int
	row.Scan(&tag)
	return tag
}
func GetTag(db *sql.DB, tagID int) {
	rows, _ := db.Query("SELECT node_id FROM node_tags WHERE tag_id = ?", tagID)
	defer rows.Close()

	var nodes []int
	for rows.Next() {
		var node int
		if err := rows.Scan(&node); err != nil {
		}
		nodes = append(nodes, node)
	}
	for _, node := range nodes {
		nodey := GetNode(db, node)
		fmt.Println(nodey.Name)
	}
}

func DeleteTag() {
}
