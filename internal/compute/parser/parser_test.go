package parser

import (
	"niffler/internal/common"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *common.Command
		err      error
	}{
		{
			name:  "valid GET command",
			input: "GET key",
			expected: &common.Command{
				Name: common.CommandType("GET"),
				Args: []string{"key"},
			},
			err: nil,
		},
		{
			name:  "valid SET command",
			input: "SET key value",
			expected: &common.Command{
				Name: common.CommandType("SET"),
				Args: []string{"key", "value"},
			},
			err: nil,
		},
		{
			name:  "valid DEL command",
			input: "DEL key",
			expected: &common.Command{
				Name: common.CommandType("DEL"),
				Args: []string{"key"},
			},
			err: nil,
		},
		{
			name:     "empty command",
			input:    "",
			expected: nil,
			err:      errInvalidCommand,
		},
		{
			name:     "command with only name",
			input:    "GET",
			expected: nil,
			err:      errInvalidCommand,
		},
		{
			name:     "command in lowercase",
			input:    "del",
			expected: nil,
			err:      errInvalidCommand,
		},
		{
			name:     "command with characters",
			input:    "SET key_/* value$",
			expected: nil,
			err:      errInvalidFormat,
		},
		{
			name:  "command with many spaces",
			input: "SET key       value    ",
			expected: &common.Command{
				Name: common.CommandType("SET"),
				Args: []string{"key", "value"},
			},
			err: nil,
		},
		{
			name:     "command with special characters",
			input:    "SET key@value",
			expected: nil,
			err:      errInvalidFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)

			if tt.err != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			if tt.expected != nil {
				require.NotNil(t, result)
				require.Equal(t, tt.expected.Name, result.Name)
				require.Equal(t, tt.expected.Args, result.Args)
			} else {
				require.Nil(t, result)
			}
		})
	}
}
