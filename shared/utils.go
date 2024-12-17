package shared

import (
	"encoding/csv"
	"os"
)

type Utils interface {
	GetEnv(key string, defaultValue string) string
	WriteCSV(data []string) error
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func WriteCSV(data []string) error {
	file, err := os.Create("countries.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}
