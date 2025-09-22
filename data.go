package main

import (
	"os"
	"fmt"
)

const headerRow string = "date,time,activity\n"

func initDataFile(filepath string) (*os.File, error) {
	fmt.Println("[INFO]: Initializing a data file...")

	fileExisted := FileExists(filepath)

	dataFile, err := os.OpenFile(filepath, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	if !fileExisted {
		fmt.Println("[INFO]: Writing the header row to the data file...")
		n, err := dataFile.Write([]byte(headerRow))
		if err != nil {
			return nil, err
		}
		if n < len(headerRow) {
			return nil, fmt.Errorf("Only %d bytes of the header row was written to the data file", n)
		}
	}

	return dataFile, nil
}

func AddRow(args *RequiredArgs) error {
	dataFile, err := initDataFile(args.dataFilePath)
	if err != nil {
		return err
	}

	row := []byte(args.String() + "\n")
	n, err := dataFile.Write(row)
	if err != nil {
		return err
	}
	if n < len(row) {
		return fmt.Errorf("Only %d bytes of the row data was written to the data file", n)
	}

	return nil
}

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}
