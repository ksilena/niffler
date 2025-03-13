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
// Если ключ не существует, возвращает errDataNotFound.
func (d *DB) Get(key string) (string, error) {
	value, ok := d.data[key]
	if !ok {
		return "", errDataNotFound
	}

	return value, nil
}

// Set устанавливает значение для указанного ключа.
// Если ключ пустой, возвращает errKeyIsEmpty.
func (d *DB) Set(key, value string) error {
	if key == "" {
		return errKeyIsEmpty
	}

	d.data[key] = value

	return nil
}

// Del удаляет значение по указанному ключу.
func (d *DB) Del(key string) {
	delete(d.data, key)
}
