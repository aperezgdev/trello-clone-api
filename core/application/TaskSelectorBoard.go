package application

import "github.com/aperezgdev/trello-clone-api/core/domain"

type TaskFinderBoard struct {
	repository domain.TaskRepository
}

func (tf *TaskFinderBoard) Run(idBoard uint32) []domain.Task {
	return tf.repository.FindByBoard(idBoard)
}

func NewTaskFinderByBoard(repository domain.TaskRepository) *TaskFinderBoard {
	return &TaskFinderBoard{repository}
}
