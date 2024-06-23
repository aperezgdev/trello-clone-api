package application

import (
	"github.com/aperezgdev/trello-clone-api/core/domain"
)

type StatusBoardFinderUser struct {
	repository domain.StatusBoardRepository
}

func NewStatusBoardFinderUser(repository domain.StatusBoardRepository) *StatusBoardFinderUser {
	return &StatusBoardFinderUser{
		repository,
	}
}

func (sbf *StatusBoardFinderUser) Run(idUser uint32) []domain.StatusBoard {
	return sbf.repository.FindByUser(idUser)
}
