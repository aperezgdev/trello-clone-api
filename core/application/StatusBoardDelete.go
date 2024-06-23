package application

import "github.com/aperezgdev/trello-clone-api/core/domain"

type StatusBoardDelete struct {
	repository domain.StatusBoardRepository
}

func NewStatusBoardDelete(repository domain.StatusBoardRepository) *StatusBoardDelete {
	return &StatusBoardDelete{
		repository,
	}
}

func (sbd *StatusBoardDelete) Run(id uint32) {
	sbd.repository.Delete(id)
}
