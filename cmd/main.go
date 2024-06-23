package main

import (
	"fmt"

	"github.com/aperezgdev/trello-clone-api/env"
	"github.com/aperezgdev/trello-clone-api/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	envVar := env.GetEnvApp(".env")

	r := gin.Default()

	fmt.Println(envVar)

	routes.InitCustomRoutes(r)
	routes.InitStatusBoardRoutes(r)
	routes.InitTaskRoutes(r)

	r.Run(":" + envVar.API_PORT)
}
