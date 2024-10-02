package csvservice

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func ToSlice(filename string) ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fullPath := filepath.Join(homeDir, ".oceanspy", "whatwords", filename)

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	words := []string{}

	// Assuming each word is in a separate row
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file or error
		}
		if len(record) > 0 {
			words = append(words, record[0]) // Take the first column (single word per line)
		}
	}

	return words, nil
}

func ToMap(filename string) (map[string]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fullPath := filepath.Join(homeDir, ".oceanspy", "whatwords", filename)

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	dataMap := make(map[string]string)

	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file or error
		}
		if len(record) >= 2 {
			dataMap[record[0]] = record[1] // First column is the key, second is the value
		}
	}

	return dataMap, nil
}
