package windirection

import (
	"errors"
	"math"
)

func DegreesToCardinal(degrees float32, points int8, lang string) (*Cardinal, error) {
	if degrees > 360 || degrees < 0 {
		return nil, errors.New("invalid degrees value")
	}
	if points != 16 && points != 8 && points != 4 {
		return nil, errors.New("invalid points value")
	}
	cardinals := getLangCardinals(lang)

	steps := 360 / int(points)
	index := (math.Trunc((float64(degrees)+float64(steps)/2)/float64(steps)) * 16) / float64(points)
	idx := int(index)
	if idx > 15 {
		idx = 0
	}
	return &cardinals[idx], nil
}
