package last

import (
	"math"
	"time"
)


// Years return years from startYear to current year
func Years(startYear int) int {
	return time.Now().Year() - startYear + 1
}

// Seasons return seasons from startYear to current year
func Seasons(startYear, startSeason int) int {
	year, month, _ := time.Now().Date()
	switch int(math.Ceil(float64(int(month) / 3.0))) {
	case 1:
		return 1 + (year - startYear - 1) * 4 + startSeason
	case 2:
		return 2 + (year - startYear - 1) * 4 + startSeason
	case 3:
		return 3 + (year - startYear - 1) * 4 + startSeason
	case 4:
		return 4 + (year - startYear - 1) * 4 + startSeason
	}
	return (year - startYear - 1) * 4 + startSeason
}

// Months return months from startYear to current year
func Months(startYear, startMonth int) int {
	year, month, _ := time.Now().Date()
	return int(month) + (year - startYear - 1) * 12 + 13 - startMonth
}
