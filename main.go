package main

import (
	"Back-end/database"
	"Back-end/router"
	"Back-end/utils"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	router.Init(r)

	err := r.Run(":9090")
	if err != nil {
		utils.LogError(err)
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
