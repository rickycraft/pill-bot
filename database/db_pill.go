package database

import (
	"database/sql"
	"math"
	"time"
)

const (
	day1        = 24 * time.Hour
	hours28days = 28 * day1 // 28days in hours
	hours21days = 21 * day1 // 21days in hours
	firstHour   = 9
)

func IsFirstDay(db *sql.DB, dt time.Time) bool {
	nextFirstDay := GetFirstPillDay(db).Add(hours28days)
	diffDay := AbsDiffDays(nextFirstDay, dt)
	return diffDay%28 == 0
}

func GetCurrentDay() time.Time {
	return getCurrentDay(time.Now())
}

func getCurrentDay(dt time.Time) time.Time {
	if dt.Hour() >= firstHour {
		return dt
	} else {
		return dt.Add(-14 * time.Hour)
	}
}

func IsWhite(db *sql.DB, _dt time.Time) bool {
	dt := _dt.UTC()
	first := GetFirstPillDay(db).UTC()
	last := first.Add(hours21days).UTC()
	return (dt.After(first) && dt.Before(last)) || dt.Equal(first)
}

func IsGreen(db *sql.DB, dt time.Time) bool {
	return !IsWhite(db, dt)
}

func AbsDiffDays(dt1 time.Time, dt2 time.Time) int64 {
	return int64(math.Abs(diffDays(dt1, dt2)))
}

func DiffDays(dt1 time.Time, dt2 time.Time) int64 {
	return int64(diffDays(dt1, dt2))
}

func diffDays(dt1 time.Time, dt2 time.Time) float64 {
	diffHour := diffHours(dt1, dt2)
	return diffHour / 24
}

func diffHours(dt1 time.Time, dt2 time.Time) float64 {
	return dt1.UTC().Sub(dt2.UTC()).Hours()
}
