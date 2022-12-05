package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      bool
	Priority    bool
	Description string
	CreatedAt   time.Time
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345"
	dbName := "gotest"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	return db
}

type Repository interface {
	Find(id int) (Task, error)
	Add(task Task) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Find(id int) (Task, error) {
	var task Task = Task{}

	rows, err := r.db.Query("SELECT * FROM tasks WHERE id=?", id)
	defer rows.Close()
	if err != nil {
		return task, err
	}

	if rows.Next() == false {
		return task, errors.New("row next error")
	}

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.StartDate, &task.DueDate, &task.Status, &task.Priority, &task.Description, &task.CreatedAt)
		if err != nil {
			return task, errors.New("error iterating through rows")
		}
	}

	if err := rows.Err(); err != nil {
		return task, err
	}

	return task, nil
}

func (r repository) Add(task Task) (int64, error) {
	fmt.Println("passed 0")

	stmt, err := r.db.Prepare("INSERT INTO tasks (title,start_date,due_date,status,priority,description,created_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	fmt.Println("passed 1")

	res, err := stmt.Exec(task.Title, task.StartDate, task.DueDate, task.Status, task.Priority, task.Description, task.CreatedAt)
	if err != nil {
		return 0, err
	}

	fmt.Println("passed 2")

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println("passed 3")

	return lastID, nil
}

func main() {
	db := dbConn()
	myDB := NewRepository(db)
	// var task Task

	// 1 nolu kayıdı bul
	/* if task, err = myDB.Find(1); err != nil {
		log.Println(err)
		return
	}

	jsonResult, _ := json.MarshalIndent(task, "", "    ")
	fmt.Println(string(jsonResult))
	*/

	task := Task{
		Title:       "new task",
		StartDate:   time.Now(),
		DueDate:     time.Now(),
		Status:      true,
		Priority:    true,
		Description: "example description",
		CreatedAt:   time.Now(),
	}

	lastID, err := myDB.Add(task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lastID)
}
