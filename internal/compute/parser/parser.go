package parser

import (
	"errors"
	"niffler/internal/common"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^\w+$`)

var (
	errInvalidCommand = errors.New("invalid command")
	errInvalidFormat  = errors.New("invalid format")
)

// Parse преобразует строковую команду в структуру domain.Command.
// Разбивает входную строку на отдельные части и проверяет каждую на соответствие
// регулярному выражению. После чего формирует структуру команды, состоящую из
// имени команды и её аргументов.
func Parse(cmd string) (*common.Command, error) {
	parts := strings.Fields(cmd)
	if len(parts) <= 1 {
		return nil, errInvalidCommand
	}

	for _, part := range parts {
		if !re.MatchString(part) {
			return nil, errInvalidFormat
		}
	}

	return &common.Command{
		Name: common.CommandType(parts[0]),
		Args: parts[1:],
	}, nil
}
