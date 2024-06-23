package domain

type TaskRepository interface {
	Create(t Task)
	Delete(id uint32)
	Update(t Task)
	FindByBoard(idBoard uint32) []Task
}
