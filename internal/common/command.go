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
	errUnknownCmdType = errors.New("unknown command type")

	argsCount = map[CommandType]int{
		GetCommand: 1,
		SetCommand: 2,
		DelCommand: 1,
	}
)

// Ошибки, используемые в валидации команд.
type Command struct {
	Name CommandType
	Args []string
}

// Valid проверяет корректность команды.
// Возвращает true, если команда валидна, и ошибку в противном случае.
func (c *Command) Valid() (bool, error) {
	cnt, ok := argsCount[c.Name]
	if !ok {
		return false, errUnknownCmdType
	}

	if len(c.Args) != cnt {
		return false, errWrongArgsCount
	}

	return true, nil
}
