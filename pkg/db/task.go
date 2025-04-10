package db

import (
	"database/sql"
	"net/http"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

var DB *sql.DB

func AddTask(task *Task) (int64, error) {
	var id int64

	query := `
		INSERT INTO scheduler (date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	return id, err
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {

}
