package handlers

import (
	"net/http"
	"strconv"

	"github.com/aperezgdev/trello-clone-api/core/application"
	"github.com/aperezgdev/trello-clone-api/core/domain"
	"github.com/aperezgdev/trello-clone-api/core/infrastructure"
	"github.com/aperezgdev/trello-clone-api/internal/database"
	"github.com/gin-gonic/gin"
)

type CustomHandler struct {
	statusBoardFinderUser application.StatusBoardFinderUser
	taskFinderBoard       application.TaskFinderBoard
}

func NewCustomHandler() *CustomHandler {
	client, err := database.InitMongoClient()

	if err != nil {
		panic(err)
	}

	statusBoardRepository := infrastructure.NewStatusBoardMongoRepository(client)
	taskRepository := infrastructure.NewTaskMongoRepository(client)

	return &CustomHandler{
		statusBoardFinderUser: *application.NewStatusBoardFinderUser(statusBoardRepository),
		taskFinderBoard:       *application.NewTaskFinderByBoard(taskRepository),
	}
}

type DataFromUser struct {
	Id    uint32        `json:"id"`
	Title string        `json:"title"`
	Tasks []domain.Task `json:"tasks"`
}

func (ch *CustomHandler) DataFromUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		panic(err)
	}

	boards := ch.statusBoardFinderUser.Run(uint32(id))
	var data []DataFromUser = make([]DataFromUser, 0)

	for i := 0; i < len(boards); i++ {
		board := boards[i]

		tasksBoard := ch.taskFinderBoard.Run(board.Id)

		data = append(data, DataFromUser{
			Id:    board.Id,
			Title: board.Title,
			Tasks: tasksBoard,
		})
	}

	if cap(data) == 0 {
		ctx.IndentedJSON(http.StatusOK, boards)
	} else {
		ctx.IndentedJSON(http.StatusOK, data)
	}

}
