package utils

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func GenerateRandomFilename(originalName string) (string, error) {
	ext := filepath.Ext(originalName)

	timestamp := time.Now().Format("20060102_150405")

	return fmt.Sprintf("%s_%s%s", timestamp, uuid.New().String()[:8], ext), nil
}
