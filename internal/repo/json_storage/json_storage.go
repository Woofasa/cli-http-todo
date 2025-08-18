package json_storage

import (
	"encoding/json"
	"errors"
	"main/internal/domain"
	"main/pkg"
	"os"
)

type JSONStorage struct{}

func (JSONStorage) Save(tasks *domain.TaskList) error {
	if !pkg.FileExists("./internal/repo/json_storage/storage.json") {
		os.Create("./internal/repo/json_storage/storage.json")
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return errors.New("encoding error")
	}
	if err := os.WriteFile("./internal/repo/json_storage/storage.json", data, 0644); err != nil {
		return errors.New("writing error")
	}
	return nil
}

func (JSONStorage) Load(tasks *domain.TaskList) error {
	reader, err := os.ReadFile("./internal/repo/json_storage/storage.json")
	if err != nil {
		return errors.New("reading error")
	}
	if err := json.Unmarshal(reader, tasks); err != nil {
		return errors.New("decoding error")
	}
	return nil
}
