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

type TaskHandler struct {
	creator application.TaskCreator
	delete  application.TaskDelete
}

func NewTaskHandler() *TaskHandler {
	client, err := database.InitMongoClient()

	if err != nil {
		panic(err)
	}

	repository := infrastructure.NewTaskMongoRepository(client)

	return &TaskHandler{
		creator: *application.NewTaskCreator(repository),
		delete:  *application.NewTaskDelete(repository),
	}
}

func (th *TaskHandler) CreateTask(ctx *gin.Context) {
	task := domain.Task{}

	err := ctx.ShouldBind(&task)

	if err == nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	th.creator.Run(task)
	ctx.Status(http.StatusCreated)
}

func (th *TaskHandler) DeleteTask(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err == nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	th.delete.Run(uint32(id))
	ctx.Status(http.StatusOK)
}
