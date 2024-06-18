// @title Phil
// @description phil is a dumbass
// @version 0.1
// @host localhost:9000
// @BasePath /
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Prepare for graceful shutdown
	go func() {
		// Channel to listen for termination signals
		stop := make(chan os.Signal, 1)
		// Catch SIGTERM and SIGINT
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

		// Block until a signal is received
		sig := <-stop
		log.Info().Msgf("Received signal: %s", sig)

		// Generate a random duration between 10 and 30 seconds
		randomDelay := rand.Intn(21) + 10 // rand.Intn(21) generates a number from 0 to 20
		delayDuration := time.Duration(randomDelay) * time.Second

		// Log the duration
		log.Info().Msgf("Setting a random shutdown delay of %v seconds", randomDelay)

		time.Sleep(delayDuration)

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		log.Info().Msg("stopping server")
		server.Stop()
	}()

	// run Server
	err = server.Start(9000)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("server error: %v", err))
	}
}
