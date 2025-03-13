package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"niffler/internal/compute/parser"
	"niffler/internal/service"
	"niffler/internal/storage/engine"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		slog.Error("create logger", "err", err)
		os.Exit(1)
	}
	defer logger.Sync()

	db := engine.New()
	parser := parser.New()

	svc := service.New(logger, db, parser)

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
