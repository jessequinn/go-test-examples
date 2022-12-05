package main

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFind(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"title", "start_date", "due_date", "status", "priority", "description", "created_at"}).
		AddRow("title", time.Now(), time.Now(), 1, 1, "description", time.Now())

	mock.ExpectQuery("SELECT * FROM tasks").WithArgs(1).WillReturnRows(rows)

	myDB := NewRepository(db)

	// run the code with the database mock
	if _, err := myDB.Find(1); err != nil {
		t.Errorf("something went wrong: %s", err.Error())
	}
}

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var id int64

	task := Task{
		Title:       "title",
		StartDate:   time.Now(),
		DueDate:     time.Now(),
		Status:      true,
		Priority:    true,
		Description: "description",
		CreatedAt:   time.Now(),
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO tasks (title,start_date,due_date,status,priority,description,created_at) VALUES($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(task.Title, task.StartDate, task.DueDate, task.Status, task.Priority, task.Description, task.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

	myDB := NewRepository(db)

	lastID, err := myDB.Add(task)
	if err != nil {
		t.Errorf("something went wrong: %s", err.Error())
	}

	fmt.Println(lastID)
}
