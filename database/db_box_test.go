package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapBox(t *testing.T) {
	db := Init()
	defer db.Close()

	lastBox := GetLastBox(db)
	assert.NotNil(t, lastBox)
	box := MapBox(lastBox)
	assert.NotNil(t, box)
	count := 0
	for idx := range box {
		if box.IsEmpty(idx) {
			count++
		}
		t.Log(box[idx])
	}
	assert.Equal(t, len(lastBox), count)
}

func TestTake(t *testing.T) {
	db := Init()
	defer db.Close()

	box := MapBox(GetLastBox(db))
	assert.NotNil(t, box)
	for idx := range box {
		if box.IsEmpty(idx) {
			assert.NotNil(t, box.Take(idx))
		} else {
			assert.Nil(t, box.Take(idx))
		}
	}
	for idx := range box {
		assert.True(t, box.IsEmpty(idx))
		assert.False(t, box.IsFull(idx))
	}
}

func TestColor(t *testing.T) {
	db := Init()
	defer db.Close()

	box := MapBox(GetLastBox(db))
	assert.NotNil(t, box)

	assert.Equal(t, white, GetColor(1))
	assert.Equal(t, white, GetColor(20))
	assert.Equal(t, green, GetColor(21))
	assert.Equal(t, green, GetColor(27))
}
