package domain

type StatusBoard struct {
	Id     uint32 `json:"id"`
	Title  string `json:"title"`
	UserId uint32 `json:"userId"`
}
