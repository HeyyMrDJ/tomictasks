package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type List struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func GetListIDByName(db *sql.DB, name string) int {
	var id int
	query := `SELECT id from lists WHERE name = ?`
	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("list with name '%s' not found", name)
		}
		return 0
	}
	return id
}

func GetListNameByID(db *sql.DB, id int) string {
	var name string
	query := `SELECT name from lists WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("list with id '%d' not found", id)
		}
		return ""
	}
	return name
}

func CreateList(db *sql.DB, name string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := tx.Prepare("insert into lists(name) values(?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetLists(db *sql.DB, scope string) []List {
	var query string
	if scope == "" {
		query = fmt.Sprintf("select id, name from lists")
	} else {
		query = fmt.Sprintf("select id, name from lists where list_id = ?")
	}
	rows, err := db.Query(query, scope)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	lists := []List{}
	for rows.Next() {
		var list List
		err = rows.Scan(&list.ID, &list.Name)
		if err != nil {
			log.Fatal(err)
		}
		lists = append(lists, list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return lists
}

func GetList(db *sql.DB, listID int) []Task {
	var query string
	query = fmt.Sprintf("select id, title, list_id, due_date from tasks where list_id = ?")
	rows, err := db.Query(query, listID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tasks := []Task{}
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Title, &task.ListID, &task.DueDate)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return tasks
}

func DeleteList(db *sql.DB, id int) {
	stmt, err := db.Prepare("DELETE FROM lists where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}

func UpdateList(db *sql.DB, id int, name string) {
	list := List{ID: id, Name: name}
	stmt, err := db.Prepare("UPDATE lists set name = ? where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(list.Name, list.ID)
	if err != nil {
		fmt.Println(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}
