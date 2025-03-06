package common

import "errors"

// CommandType представляет тип команды в виде строкового значения.
type CommandType string

// Константы для поддерживаемых типов команд.
const (
	GetCommand CommandType = "GET"
	SetCommand CommandType = "SET"
	DelCommand CommandType = "DEL"
)

var (
	errWrongArgsCount = errors.New("invalid args count")
)

// Ошибки, используемые в валидации команд.
type Command struct {
	Name CommandType
	Args []string
}

// Valid проверяет корректность команды:
// для SET требуется ровно 2 аргумента;
// для остальных команд требуется 1 аргумент.
// Возвращает true, если команда валидна, и ошибку в противном случае.
func (c *Command) Valid() (bool, error) {
	argsLen := 1
	if c.Name == SetCommand {
		argsLen = 2
	}

	if len(c.Args) != argsLen {
		return false, errWrongArgsCount
	}

	return true, nil
}
