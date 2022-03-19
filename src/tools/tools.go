package tools

import (
	"fmt"
	"time"
)

// FormatVMDate returns parsed time from VM API
func FormatVMDate(date string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02T15:04:05-0700", date) //Not perfect format, but still works.
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}

// ByteCountIEC returns formatted byte count. Grabbed from https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/.
func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
