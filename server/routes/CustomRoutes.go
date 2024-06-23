package routes

import (
	"github.com/aperezgdev/trello-clone-api/server/handlers"
	"github.com/gin-gonic/gin"
)

var customHandler = handlers.NewCustomHandler()

func InitCustomRoutes(router *gin.Engine) {
	router.POST("/getData/users/:id", customHandler.DataFromUser)
}
