package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)
type task struct{
Data string
}

// TODO: if path is absent, return an empty state with NextID=1.
// If the file exists, decode exactly one JSON document, validate it via Restore,
// and return a meaningful error for corrupted data. Do NOT silently reset it.
func Load(path string) (task, error) {

	data, err := os.ReadFile(path) 
	if errors.Is(err, os.ErrNotExist){
		return task{Data: "0"}, nil
	}

	decJson := json.NewDecoder(bytes.NewReader(data))
	decJson.DisallowUnknownFields()
	
	var s task
	if err := decJson.Decode(&s); err != nil {
    	fmt.Print(err)
	}

	return task{Data: "123"}, nil
}

func main(){
	Load("fl.json")
}

// TODO: validate the state, marshal it using encoding/json, save to a
// temporary file in the same directory and replace the target using os.Rename.
// Do not remove the old file before a successful temporary write.
// func Save(path string, state tasks.State) error {
// 	return nil
// }
