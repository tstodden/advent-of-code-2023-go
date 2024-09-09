package day1

import (
	"errors"
	"strconv"
	"strings"
)

var ErrColorFormat = errors.New("color must have exactly 2 parts")

var ErrInvalidColor = errors.New("invalid color present")

type Sample struct {
	Green int
	Blue  int
	Red   int
}

func NewSample(s string) (Sample, error) {
	var result Sample

	colors := strings.Split(s, ", ")
	for _, v := range colors {
		parts := strings.Split(v, " ")
		if len(parts) != 2 {
			return result, ErrColorFormat
		}

		count, err := strconv.Atoi(parts[0])
		if err != nil {
			return result, err
		}

		switch parts[1] {
		case "blue":
			result.Blue += count
		case "green":
			result.Green += count
		case "red":
			result.Red += count
		default:
			return result, ErrInvalidColor
		}
	}

	return result, nil
}
