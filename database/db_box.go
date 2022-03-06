package database

import "fmt"

const (
	boxSize = 28
	white   = true
	green   = false
	full    = false
	empty   = true
)

type Box [boxSize]bool

func MapBox(pillDays []PillDay) *Box {
	if len(pillDays) > boxSize {
		return nil
	}
	var box Box
	for idx := range pillDays {
		box[idx] = true
	}
	return &box
}

func GetColor(day int) bool {
	return day < 21
}

func (bx *Box) IsFull(day int) bool {
	return !bx[day]
}

func (bx *Box) IsEmpty(day int) bool {
	return bx[day]
}

func (bx *Box) Take(day int) error {
	if bx.IsEmpty(day) {
		return fmt.Errorf("boxDay is already empty")
	}
	bx[day] = empty
	return nil
}
