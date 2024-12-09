package utils

import (
	"log/slog"
	"os"
	"strings"
)

func getInput() ([]string, error) {
	fileData, err := os.ReadFile("input.txt")
	if err != nil {
		slog.Error("error", "e", err.Error())
		return []string{}, err
	}
	file := string(fileData)
	lines := strings.Split(file, "\n")

	return lines, nil
}
