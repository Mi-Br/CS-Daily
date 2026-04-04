package main

import (
	"encoding/json"
	"errors"
	"os"
)

const FILE_NAME = "storage.json"

type Store struct {
	File os.File
}

func (s *Store) Load(b *Backlog) error {
	data, err := os.ReadFile(FILE_NAME)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.Create(FILE_NAME)
			data, err := json.Marshal(b)
			if err != nil {
				return err
			}
			os.WriteFile(FILE_NAME, data, 0644)
			return nil
		}
		return err
	}
	err = json.Unmarshal(data, b)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) Save(b *Backlog) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return os.WriteFile(FILE_NAME, data, 0644)
}
