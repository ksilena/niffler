package engine

import "errors"

var (
	errDataNotFound = errors.New("data not found")
	errKeyIsEmpty   = errors.New("key is empty")
)

// DB представляет in-memory key-value базу данных.
type DB struct {
	data map[string]string
}

// New создает и возвращает новую экземпляр базы данных DB.
func New() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

// Get возвращает значение по указанному ключу.
// Если ключ пустой, возвращает ErrKeyIsEmpty.
// Если ключ не существует, возвращает ErrDataNotFound.
func (d *DB) Get(key string) (string, error) {
	if key == "" {
		return "", errKeyIsEmpty
	}

	value := d.data[key]
	if value == "" {
		return "", errDataNotFound
	}

	return value, nil
}

// Set устанавливает значение для указанного ключа.
// Если ключ пустой, возвращает ErrKeyIsEmpty.
func (d *DB) Set(key, value string) error {
	if key == "" {
		return errKeyIsEmpty
	}

	d.data[key] = value

	return nil
}

// Del удаляет значение по указанному ключу.
// Если ключ пустой, возвращает ErrKeyIsEmpty.
func (d *DB) Del(key string) error {
	if key == "" {
		return errKeyIsEmpty
	}

	delete(d.data, key)

	return nil
}
