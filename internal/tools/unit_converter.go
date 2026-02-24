package tools

import (
	"fmt"
	"math"
)

func FormatBytes(bytes int64) string {
	if bytes <= 1023 {
		return fmt.Sprintf("%d B", bytes)
	}

	units := []string{"B", "KB", "MB", "GB", "TB"}
	unitIndex := 0
	value := float64(bytes)

	for value >= 100 && unitIndex < len(units)-1 {
		unitIndex++
		value = value / 1024
	}

	if value < 1 && unitIndex > 0 {
		unitIndex--
		value = value * 1024
	}

	var result string
	if value == math.Trunc(value) {
		result = fmt.Sprintf("%.0f", value)
	} else if value < 10 {
		result = fmt.Sprintf("%.2f", value)
	} else {
		result = fmt.Sprintf("%.1f", value)
	}

	return result + " " + units[unitIndex]
}
