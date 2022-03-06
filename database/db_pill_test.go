package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	first = time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC)
)

func TestAbsDiffDays(t *testing.T) {
	dates := []time.Time{
		time.Date(2022, 2, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 6, 10, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 28, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 28, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 29, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 29, 12, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 29, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 30, 10, 0, 0, 0, time.UTC),
	}
	diffs := []int64{
		18, 0, 0, 5, 27, 27, 28, 28, 28, 29, 29,
	}
	for idx, dt := range dates {
		diff := AbsDiffDays(dt, first)
		t.Logf("absdiff %d, expected %d", diff, diffs[idx])
		assert.Equal(t, diff, diffs[idx])
		// opposite dates should be equal
		assert.Equal(t, diff, AbsDiffDays(first, dt))
	}
}

func TestDiffDays(t *testing.T) {
	dates := []time.Time{
		time.Date(2022, 2, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 6, 10, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 28, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 28, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 29, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 29, 12, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 29, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 30, 10, 0, 0, 0, time.UTC),
	}
	diffs := []int64{
		-18, 0, 0, 5, 27, 27, 28, 28, 28, 29, 29,
	}
	for idx, dt := range dates {
		diff := DiffDays(dt, first)
		t.Logf("diff %d, expected %d", diff, diffs[idx])
		assert.Equal(t, diff, diffs[idx])
		// opposite dates should be inverse
		assert.Equal(t, diff, -1*DiffDays(first, dt))
	}
}

func TestIsWhite(t *testing.T) {
	db := Init()
	defer db.Close()

	dates := []time.Time{
		time.Date(2022, 2, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 6, 10, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 20, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 21, 12, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 21, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 21, 23, 59, 0, 0, time.UTC),
		time.Date(2022, 3, 22, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 22, 10, 0, 0, 0, time.UTC),
	}
	diffs := []bool{
		false, true, true, true, true, true, true, true, true, true, false, false,
	}
	assert.Equal(t, len(dates), len(diffs))
	for idx, dt := range dates {
		res := IsWhite(db, dt)
		t.Logf("dt %s, res %t, expected %t", dt.Format(time.RFC3339Nano), res, diffs[idx])
		assert.Equal(t, diffs[idx], res)
	}
}

func TestIsGreen(t *testing.T) {
	db := Init()
	defer db.Close()

	dates := []time.Time{
		time.Date(2022, 2, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 6, 10, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 20, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 21, 12, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 21, 23, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 21, 23, 59, 0, 0, time.UTC),
		time.Date(2022, 3, 22, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 22, 10, 0, 0, 0, time.UTC),
	}
	diffs := []bool{
		false, true, true, true, true, true, true, true, true, true, false, false,
	}
	assert.Equal(t, len(dates), len(diffs))
	for idx, dt := range dates {
		res := IsGreen(db, dt)
		t.Logf("dt %s, res %t, expected %t", dt.Format(time.RFC3339Nano), res, !diffs[idx])
		assert.Equal(t, res, !diffs[idx])
	}
}

func TestGetCurrentDay(t *testing.T) {
	dates := []time.Time{
		time.Date(2022, 3, 10, 14, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 10, 21, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 11, 2, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 11, 9, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 11, 10, 0, 0, 0, time.UTC),
	}
	res := []int{
		10,
		10,
		10,
		11,
		11,
	}

	for idx, dt := range dates {
		curr := getCurrentDay(dt)
		t.Logf("%s %s", dt, curr)
		assert.Equal(t, res[idx], curr.Day())
	}

}

func TestIsFirstDay(t *testing.T) {
	db := Init()
	defer db.Close()

	assert.True(t, IsFirstDay(db, first))
	assert.True(t, IsFirstDay(db, time.Date(2022, 3, 29, 0, 0, 0, 0, time.UTC)))
	assert.False(t, IsFirstDay(db, time.Date(2022, 3, 28, 0, 0, 0, 0, time.UTC)))
	assert.False(t, IsFirstDay(db, time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC)))
	assert.False(t, IsFirstDay(db, time.Date(2022, 2, 10, 0, 0, 0, 0, time.UTC)))
}

func TestGeneric(t *testing.T) {
	db := Init()
	defer db.Close()

}
