package migration

import (
	"os"
	"strings"
)

func getStatementsFromFile(str string) []string {
	statements := strings.Split(str, ";")
	array := make([]string, 0)
	for _, statement := range statements {
		if strings.TrimSpace(statement) != "" {
			array = append(array, statement)
		}
	}
	return array
}

func readSqlFile(path string) (string, error) {
	filePath := "./" + path

	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	str := string(b)
	return str, nil
}

func readSqlFiles(paths []string) ([]string, error) {
	results := make([]string, 0)
	for _, path := range paths {
		str, err := readSqlFile(path)
		if err != nil {
			return nil, err
		}

		statements := getStatementsFromFile(str)
		results = append(results, statements...)
	}
	return results, nil
}
