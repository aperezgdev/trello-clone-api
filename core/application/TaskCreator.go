package application

import "github.com/aperezgdev/trello-clone-api/core/domain"

type TaskCreator struct {
	repository domain.TaskRepository
}

func NewTaskCreator(repository domain.TaskRepository) *TaskCreator {
	return &TaskCreator{
		repository,
	}
}

func (tc *TaskCreator) Run(t domain.Task) {
	tc.repository.Create(t)
}
