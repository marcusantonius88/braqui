package reminder

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseDate(text string, now time.Time) (time.Time, error) {
	text = strings.TrimSpace(strings.ToLower(text))

	if text == "hoje" {
		return now.Truncate(24 * time.Hour), nil
	}

	if text == "amanhã" || text == "amanha" {
		return now.Truncate(24*time.Hour).Add(24 * time.Hour), nil
	}

	if strings.HasPrefix(text, "daqui a ") && strings.HasSuffix(text, " dias") {
		middle := strings.TrimPrefix(text, "daqui a ")
		middle = strings.TrimSuffix(middle, " dias")
		n, err := strconv.Atoi(strings.TrimSpace(middle))
		if err == nil && n > 0 {
			return now.Truncate(24*time.Hour).Add(time.Duration(n) * 24 * time.Hour), nil
		}
	}

	if strings.HasPrefix(text, "em ") && strings.HasSuffix(text, " dias") {
		middle := strings.TrimPrefix(text, "em ")
		middle = strings.TrimSuffix(middle, " dias")
		n, err := strconv.Atoi(strings.TrimSpace(middle))
		if err == nil && n > 0 {
			return now.Truncate(24*time.Hour).Add(time.Duration(n) * 24 * time.Hour), nil
		}
	}

	if strings.HasPrefix(text, "dia ") {
		dayStr := strings.TrimPrefix(text, "dia ")
		n, err := strconv.Atoi(strings.TrimSpace(dayStr))
		if err == nil && n >= 1 && n <= 31 {
			year, month, _ := now.Date()
			day := n
			candidate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			if candidate.Before(now.Truncate(24 * time.Hour)) {
				candidate = time.Date(year, month+1, day, 0, 0, 0, 0, time.UTC)
			}
			return candidate, nil
		}
	}

	return time.Time{}, fmt.Errorf("unrecognized date: %s", text)
}
