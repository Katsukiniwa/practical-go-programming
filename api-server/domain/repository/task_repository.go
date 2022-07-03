package repository

import (
	"../model"
)

type TaskRepository interface {
	Store(task *model.Task) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
}
