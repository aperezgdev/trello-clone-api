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

type StatusBoardHandler struct {
	creator application.StatusBoardCreator
	delete  application.StatusBoardDelete
	finder  application.StatusBoardFinderUser
}

func NewStatusBoardHandler() *StatusBoardHandler {
	client, err := database.InitMongoClient()

	if err != nil {
		panic(err)
	}

	repository := infrastructure.NewStatusBoardMongoRepository(client)

	return &StatusBoardHandler{
		creator: *application.NewStatusBoardCreator(repository),
		delete:  *application.NewStatusBoardDelete(repository),
		finder:  *application.NewStatusBoardFinderUser(repository),
	}
}

func (sbh *StatusBoardHandler) CreateStatusBoard(ctx *gin.Context) {
	statusBoard := domain.StatusBoard{}

	err := ctx.ShouldBind(&statusBoard)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	sbh.creator.Run(statusBoard)
	ctx.Status(http.StatusCreated)
}

func (sbh *StatusBoardHandler) DeleteStatusBoard(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err == nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sbh.delete.Run(uint32(id))
	ctx.Status(http.StatusOK)
}

func (sbh *StatusBoardHandler) FinderByUserStatusBoard(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	statusBoards := sbh.finder.Run(uint32(id))

	ctx.JSON(http.StatusOK, statusBoards)
}
