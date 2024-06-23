package domain

type Task struct {
	Id            uint32 `json:"id"`
	Title         string `json:"title"`
	Completed     bool   `json:"completed"`
	BoardStatusId uint32 `json:"boardStatusId"`
}
