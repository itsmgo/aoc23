package common

import (
	"log"
	"os"
)

func LoadInputContent(filename string) string {
	body, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	return string(body)
}
