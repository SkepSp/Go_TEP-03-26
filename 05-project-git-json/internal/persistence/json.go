package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"example.com/go-practice-01/internal/tasks"
)

// TODO: if path is absent, return an empty state with NextID=1.
// If the file exists, decode exactly one JSON document, validate it via Restore,
// and return a meaningful error for corrupted data. Do NOT silently reset it.
func Load(path string) (tasks.State, error) {

	data, err := os.ReadFile(path) 
	if errors.Is(err, os.ErrNotExist){
		return tasks.State{NextID: 1}, nil //Json файлик не найден
	}

	decJson := json.NewDecoder(bytes.NewReader(data))
	decJson.DisallowUnknownFields()

	var decState tasks.State
	err = decJson.Decode(&decState)
	if err != nil {
		return tasks.State{}, fmt.Errorf("Ошибка чтения json %w", err)
	}
	var tmp tasks.State
	err = decJson.Decode(&tmp)
	if !errors.Is(err,io.EOF) {
		if (err != nil ) {
			return tasks.State{}, fmt.Errorf("Ошибка чтения json %w", err)
		}
		return tasks.State{}, fmt.Errorf("Ошибка файла json")
	}
	err = tasks.NewMemoryStore().Restore(decState)
	if err != nil {
		return tasks.State{}, fmt.Errorf("Ошибка данных в файле %w", err)
	}

	return decState, nil
}

// TODO: validate the state, marshal it using encoding/json, save to a
// temporary file in the same directory and replace the target using os.Rename.
// Do not remove the old file before a successful temporary write.
func Save(path string, state tasks.State) error {
	err := tasks.NewMemoryStore().Restore(state)
	if err != nil {
		return fmt.Errorf("Ошибка данных state %w", err)
	}
	encode, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("Ошибка данных state %w", err)
	}
	
	encode = append(encode, '\n')

	tmpFile, err := os.CreateTemp(filepath.Dir(path), ".tasks-*.tmp")
	if err != nil {
		return fmt.Errorf("Ошибка создания файла %w", err)
	}

	defer os.Remove(tmpFile.Name()); 
	defer tmpFile.Close()
	
	_, err = tmpFile.Write(encode)
	if err != nil{
		return fmt.Errorf("Ошибка записи %w", err)
	}
	err = tmpFile.Sync()
	if err != nil {
		return fmt.Errorf("Ошибка записи %w", err)
	}
	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("Ошибка закрытия файлы %w", err)
	}
	err = os.Rename(tmpFile.Name(), path)
	if err != nil {
		return fmt.Errorf("Ошибка переименовывания %w", err)
	}

	return nil
}
