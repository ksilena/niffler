package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"niffler/internal/service"
	"niffler/internal/storage/engine"
	"os"

	"go.uber.org/zap"
)

func main() {
	flag.Parse()
	logger, err := zap.NewDevelopment()
	if err != nil {
		slog.Error("create logger", "err", err)
		os.Exit(1)
	}
	defer logger.Sync()

	db := engine.New()

	svc := service.New(db, logger)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")

		scanner.Scan()
		output, err := svc.Handle(scanner.Text())
		if err != nil {
			logger.Error("handle command", zap.Error(err))
			continue
		}

		logger.Info("success", zap.String("response", output))
	}
}
