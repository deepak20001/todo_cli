package main

import (
	"encoding/json"
	"os"
)

// Storage[T any] is a generic struct, meaning it can store any type (T).
// The struct contains fileName, which is the name of the file where data will be stored.
type Storage[T any] struct {
	fileName string
}

// This function creates a new Storage instance.
// It returns a pointer to a Storage[T] object.
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{fileName: fileName}
}

// Save data to file
// json.MarshalIndent(data, "", " ") Converts data (of any type T) into formatted JSON.
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.fileName, fileData, 0644)
}

// / Load data from file
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
