package application

import (
	"github.com/aperezgdev/trello-clone-api/core/domain"
)

type StatusBoardCreator struct {
	repository domain.StatusBoardRepository
}

func NewStatusBoardCreator(repository domain.StatusBoardRepository) *StatusBoardCreator {
	return &StatusBoardCreator{
		repository,
	}
}

func (sbc *StatusBoardCreator) Run(sb domain.StatusBoard) {
	sbc.repository.Create(sb)
}
