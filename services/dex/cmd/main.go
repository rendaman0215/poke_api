package main

import (
	"github.com/rendaman0215/poke_api/services/dex/internal/handler"
	"github.com/rendaman0215/poke_api/services/dex/internal/logger"
)

func main() {
	// Initialize the logger
	logger.InitLogger()

	// Start the server
	handler.Handler()
}
