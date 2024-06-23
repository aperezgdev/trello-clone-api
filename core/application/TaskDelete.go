package application

import "github.com/aperezgdev/trello-clone-api/core/domain"

type TaskDelete struct {
	repository domain.TaskRepository
}

func NewTaskDelete(repository domain.TaskRepository) *TaskDelete {
	return &TaskDelete{
		repository,
	}
}

func (td *TaskDelete) Run(id uint32) {
	td.repository.Delete(id)
}
