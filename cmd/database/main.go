package main

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	defer logger.Sync()
}
