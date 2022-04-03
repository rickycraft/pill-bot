package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertPillToday(t *testing.T) {
	db := Init()
	defer db.Close()

	err := Take(db)
	assert.Nil(t, err)
}

func TestGetLastPill(t *testing.T) {
	db := Init()
	defer db.Close()

	pill := GetLastPill(db)
	assert.NotNil(t, pill)
	assert.NotEqual(t, pill, &PillDay{})
	t.Logf("Id %d, Date %s, IsFirstDay %t", pill.Id, pill.Date.Format(time.RFC3339Nano), pill.IsFirstDay)
}

func TestGetLastBox(t *testing.T) {
	db := Init()
	defer db.Close()

	days := GetLastBox(db)
	assert.NotEmpty(t, days)
	assert.True(t, len(days) < 29)
	t.Logf("Length %d\n", len(days))
	for _, pill := range days {
		assert.NotEqual(t, pill, PillDay{})
		t.Logf("Id %d, Date %s, IsFirstDay %t\n", pill.Id, pill.Date.Format(time.RFC3339Nano), pill.IsFirstDay)
	}
}

func TestGetFirstPillDay(t *testing.T) {
	db := Init()
	defer db.Close()

	dt := GetFirstPillDay(db)
	assert.Equal(t, dt, time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC))
}

// func TestGetLastMonth(t *testing.T) {
// 	db := Init()
// 	defer db.Close()

// 	days := GetLastMonth(db)
// 	assert.NotEmpty(t, days)
// 	assert.True(t, len(days) <= 30)
// 	t.Logf("Length %d\n", len(days))
// 	for _, pill := range days {
// 		assert.NotEqual(t, pill, PillDay{})
// 		t.Logf("Id %d, Date %s, IsFirstDay %t\n", pill.Id, pill.Date.Format(time.RFC3339Nano), pill.IsFirstDay)
// 	}
// }
