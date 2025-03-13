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
			err:      errDataNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := New()
			db.data["key"] = "value"

			result, err := db.Get(tt.key)

			if tt.err != nil {
				require.EqualError(t, tt.err, err.Error())
			} else {
				require.NoError(t, err)
			}

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

			if tt.err != nil {
				require.EqualError(t, tt.err, err.Error())
			} else {
				require.NoError(t, err)
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := New()
			db.data["key"] = "value"

			db.Del(tt.key)

			_, err := db.Get(tt.key)
			require.EqualError(t, err, errDataNotFound.Error())
		})
	}
}
