/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"database/sql"
	"github.com/heyymrdj/tomictasks/cmd"
	"github.com/heyymrdj/tomictasks/pkg/database"
)

var db *sql.DB

func main() {
	db := database.ConnectDB()
	defer db.Close()
	database.CreateTable(db)
	cmd.Execute(db)
	//database.CreateTask(db, "TEST", 0, "")
}
