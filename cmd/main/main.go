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
	log.Info().Msg("starting up")

	// random delay, simulating startup in production
	randomDelay := rand.Intn(21) + 10 // rand.Intn(21) generates a number from 0 to 20
	delayDuration := time.Duration(randomDelay) * time.Second
	log.Info().Msgf("setting a random startup delay of %v seconds", randomDelay)
	time.Sleep(delayDuration)
	log.Info().Msg("startup phase complete")

	// get configuration
	_, err := config.GetConfig()
	if err != nil {
		log.Fatal().Msgf("failed to load config: %s", err)
	}

	// prepare routes
	whoamiSvc := services.NewWhoamiSvc()
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
		log.Info().Msgf("received signal: %s", sig)

		// Generate a random duration between 10 and 30 seconds
		// this simulates the time during which existing connections will be handled to completion
		randomDelay := rand.Intn(21) + 10 // rand.Intn(21) generates a number from 0 to 20
		delayDuration := time.Duration(randomDelay) * time.Second

		// Log the duration
		log.Info().Msgf("setting a random shutdown delay of %v seconds", randomDelay)

		time.Sleep(delayDuration)

		// at this point, all open business should be completed and we can stop the server
		log.Info().Msg("stopping server")
		server.Stop()
	}()

	// run Server
	err = server.Start(9000)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("server error: %v", err))
	}
}
