package util

import (
	"fmt"
	"github.com/google/uuid"
	"path/filepath"
	"strconv"
	"time"
)

// GenerateFilename year/month/day/uuid-string.extension =>e.g (2024/04/30/078998-897-546.jgp)
func GenerateFilename(originalFilename string) string {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := fmt.Sprintf("%02d", now.Month())
	day := fmt.Sprintf("%02d", now.Day())
	random := uuid.New().String()

	extension := filepath.Ext(originalFilename)

	return fmt.Sprintf("%s/%s/%s/%s%s", year, month, day, random, extension)
}
