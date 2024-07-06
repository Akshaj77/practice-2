package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)

	if err != nil {
		// fmt.Println("Could not open file!")
		// fmt.Println(err)
		return nil, errors.New("failed to open file")
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read file")
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputPath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to conver data to JSON ")
	}

	file.Close()
	return nil

}

func New(inputpath string, outputpath string) FileManager {
	return FileManager{
		InputPath:  inputpath,
		OutputPath: outputpath,
	}
}
