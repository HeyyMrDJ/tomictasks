package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID          int
	ListID      int
	ListName    string
	Title       string
	Description string
	DueDate     string
	Completed   bool
	CreatedAt   time.Time
}

func CreateTask(db *sql.DB, name string, listID int, dueDate string) error {
	if listID == 0 {
		listID = 2
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO tasks (list_id, title, description, due_date) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(listID, name, name, dueDate)
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

func GetTaskIDByName(db *sql.DB, name string, listID int) int {
	var id int
	query := `SELECT id from tasks WHERE title = ? and list_id = ?`
	err := db.QueryRow(query, name, listID).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("task with name '%s' not found", name)
		}
		return 0
	}
	return id
}

func GetTasks(db *sql.DB, scope string) []Task {
	var query string
	if scope == "" {
		query = fmt.Sprintf("select id, title, list_id, due_date from tasks")
	} else {
		query = fmt.Sprintf("select id, title, list_id, due_date from tasks where list_id = ?")
	}
	rows, err := db.Query(query, scope)
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
		task.ListName = GetListNameByID(db, task.ListID)
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return tasks
}

func ReadTask(db *sql.DB, id int) {
	stmt, err := db.Prepare("SELECT id, title from tasks where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	var task Task

	err = stmt.QueryRow(id).Scan(&task.ID, &task.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows were returned!")
		} else {
			log.Fatal(err)
		}
		return
	}

	fmt.Printf("ID: %d, Name: %s\n", task.ID, task.Title)
}

func UpdateTask(db *sql.DB, id int, title string) {
	task := Task{ID: id, Title: title}
	stmt, err := db.Prepare("UPDATE tasks set title = ? where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(task.Title, task.ID)
	if err != nil {
		fmt.Println(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}

func DeleteTask(db *sql.DB, id int) bool {
	stmt, err := db.Prepare("DELETE FROM tasks where id = ?")
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
	if affected == 1 {
		return true
	} else {
		return false
	}
}
