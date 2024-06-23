package domain

type StatusBoardRepository interface {
	Create(st StatusBoard)
	Delete(id uint32)
	Update(st StatusBoard)
	FindByUser(idUser uint32) []StatusBoard
}
