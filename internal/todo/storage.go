package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{Filename: filename}
}

func (s *Storage[T]) Save(data T) error {
	bytes, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("failed to marshall data: %w", err)
	}

	err = os.WriteFile(s.Filename, bytes, 0664)

	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func (s *Storage[T]) Load() (*T, error) {
	jsonbytes, err := os.ReadFile(s.Filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var data T
	err = json.Unmarshal(jsonbytes, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall data: %w", err)
	}

	return &data, nil

}
