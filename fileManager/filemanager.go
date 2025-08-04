package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New(" Failed to Open File! ")
	}

	scanner := bufio.NewScanner(file)
	var line []string

	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to Read line in file")
	}

	file.Close()
	return line, nil
}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New(" Failed to create File.")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New(" Failed to convert data Json file.")
	}
	file.Close()
	return nil
}
