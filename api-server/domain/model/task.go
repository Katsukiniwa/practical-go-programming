package model

import (
	"errors"
)

type Task struct {
	ID      int
	Title   string
	Content string
}

func InitializeTask(title string, content string) (*Task, error) {
	if title == "" {
		return nil, errors.New("title is blank")
	}

	task := &Task{
		Title:   title,
		Content: content,
	}

	return task, nil
}

func (t *Task) Set(title string, content string) error {
	if title == "" {
		return errors.New("enter title")
	}

	t.Title = title
	t.Content = content

	return nil
}
