package main

import (
	"log"

	"github.com/Adityadangi14/centralized_logging_system/log_collector/src/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	var server *initializers.TCPServer
	server, err := initializers.NewTCPServer(":3000", 5)

	if err != nil {
		log.Fatalf("Failed to create TCP connection : %v", err)
	}

	server.Start()

	server.Wg.Wait()

}
