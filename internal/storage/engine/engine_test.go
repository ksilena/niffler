package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		expected string
		err      error
	}{
		{
			name:     "success found data",
			key:      "key",
			expected: "value",
			err:      nil,
		},
		{
			name:     "data not found",
			key:      "another",
			expected: "",
			err:      errDataNotFound,
		},
		{
			name:     "empty key",
			key:      "",
			expected: "",
			err:      errKeyIsEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := New()
			db.data["key"] = "value"

			result, err := db.Get(tt.key)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		expected string
		err      error
	}{
		{
			name:     "success set data",
			key:      "key",
			value:    "value",
			expected: "value",
			err:      nil,
		},
		{
			name:  "empty key",
			key:   "",
			value: "",
			err:   errKeyIsEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := New()

			err := db.Set(tt.key, tt.value)

			require.Equal(t, tt.err, err)

			if err == nil {
				res, err := db.Get(tt.key)
				require.Nil(t, err)
				require.Equal(t, tt.value, res)
			}
		})
	}
}

func TestDel(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		expected string
		err      error
	}{
		{
			name:     "success del data",
			key:      "key",
			expected: "value",
			err:      nil,
		},
		{
			name: "empty key",
			key:  "",
			err:  errKeyIsEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := New()
			db.data["key"] = "value"

			err := db.Del(tt.key)

			require.Equal(t, tt.err, err)

			if err == nil {
				_, err = db.Get(tt.key)
				require.Equal(t, err, errDataNotFound)
			}
		})
	}
}
