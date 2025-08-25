package todo

import (
	"errors"
	"time"

	"github.com/rodaine/table"
)

type Task struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Tasks []Task

func (tasks *Tasks) Add(title string) {
	task := Task{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*tasks = append(*tasks, task)
}

func (tasks *Tasks) ValidateIndex(index int) error {
	if index < 0 || index >= len(*tasks) {
		err := errors.New("Invalid index : index is out of bounds.")
		println(err)
		return err
	}

	return nil
}

func (tasks *Tasks) Delete(index int) error {

	t := *tasks

	if err := tasks.ValidateIndex(index); err != nil {

		return err
	}
	*tasks = append(t[:index], t[index+1:]...)

	return nil
}

func (tasks *Tasks) Toggle(index int) error {
	t := *tasks
	if err := tasks.ValidateIndex(index); err != nil {

		return err
	}
	isCompleted := t[index].Completed

	if isCompleted == false {
		completionTime := time.Now()

		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted
	*tasks = t

	return nil
}

func (tasks *Tasks) Edit(index int, newTitle string) error {

	t := *tasks
	if err := tasks.ValidateIndex(index); err != nil {

		return err
	}

	t[index].Title = newTitle

	*tasks = t

	return nil
}

func (tasks *Tasks) Print() {

	tbl := table.New("ID", "Title", "Completed", "Created At", "Completed At")

	for index, task := range *tasks {
		completed := "❌"
		completedAt := ""

		if task.Completed {
			completed = "✅"
			if task.CompletedAt != nil {
				completedAt = task.CompletedAt.Format(time.RFC1123)
			}
		}

		tbl.AddRow(index, task.Title, completed, task.CreatedAt.Format(time.RFC1123), completedAt)
	}

	tbl.Print()
}
