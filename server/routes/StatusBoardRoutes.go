package routes

import (
	"github.com/aperezgdev/trello-clone-api/server/handlers"
	"github.com/gin-gonic/gin"
)

var handler = handlers.NewStatusBoardHandler()

func InitStatusBoardRoutes(router *gin.Engine) {
	router.POST("/boards", handler.CreateStatusBoard)
	router.DELETE("/boards", handler.DeleteStatusBoard)
	router.GET("/boards/users/:id", handler.FinderByUserStatusBoard)
}
