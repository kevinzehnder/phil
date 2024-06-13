// @title Phil
// @description phil is a dumbass
// @version 0.1
// @host localhost:9000
// @BasePath /
package main

import (
	"fmt"

	"github.com/kevinzehnder/phil/pkg/handlers"
	"github.com/kevinzehnder/phil/pkg/server"
	"github.com/kevinzehnder/phil/pkg/services"

	"github.com/kevinzehnder/phil/internal/config"
	_ "github.com/kevinzehnder/phil/internal/logging"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting Up...")

	// get configuration
	_, err := config.GetConfig()
	if err != nil {
		log.Fatal().Msgf("Failed to load config: %s", err)
	}

	// create Servicer, pass in the Databaser
	whoamiSvc := services.NewWhoamiSvc()

	// create Handler, pass in the Servicer
	whoamiHandler := handlers.NewWhoamiHandler(whoamiSvc)

	// create Server. pass in the Handlers.
	server := server.NewServer(whoamiHandler)

	// run Server
	err = server.Start(9000)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to start server: %v", err))
	}
}
