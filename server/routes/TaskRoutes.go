package routes

import (
	"github.com/aperezgdev/trello-clone-api/server/handlers"
	"github.com/gin-gonic/gin"
)

var taskHandler = handlers.NewTaskHandler()

func InitTaskRoutes(router *gin.Engine) {
	router.POST("/tasks", taskHandler.CreateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
