package service

//go:generate mockery --all --output . --inpackage

import (
	"errors"
	"niffler/internal/common"

	"go.uber.org/zap"
)

// Storage определяет интерфейс для работы с хранилищем данных (установка, получение, удаление значений).
type Storage interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Del(key string)
}

type Parser interface {
	Parse(cmd string) (*common.Command, error)
}

// Service представляет сервисный слой, который обрабатывает входящие сообщения,
// используя парсер для анализа команд и хранилище для выполнения операций.
type Service struct {
	storage Storage
	parser  Parser
	logger  *zap.Logger
}

// New создает новый экземпляр сервиса с заданным хранилищем и логгером.
func New(logger *zap.Logger, storage Storage, parser Parser) *Service {
	return &Service{
		storage: storage,
		parser:  parser,
		logger:  logger,
	}
}

var (
	errUnknownCmd = errors.New("unknown command")
)

// Handle обрабатывает входящее сообщение от клиента.
func (s *Service) Handle(msg string) (string, error) {
	cmd, err := s.parser.Parse(msg)
	if err != nil {
		return "", err
	}

	if ok, err := cmd.Valid(); !ok {
		return "", err
	}

	switch cmd.Name {
	case common.GetCommand:
		return s.storage.Get(cmd.Args[0])

	case common.SetCommand:
		return "", s.storage.Set(cmd.Args[0], cmd.Args[1])

	case common.DelCommand:
		s.storage.Del(cmd.Args[0])
		return "", nil

	default:
		return "", errUnknownCmd
	}
}
