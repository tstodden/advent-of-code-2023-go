package day3

import "errors"

var ErrEmptySchematic = errors.New("schematic cannot be empty")

type Schematic struct {
	MaxX       int
	MaxY       int
	References []SchematicReference
}

func NewSchematic(lines []string) (Schematic, error) {
	result := Schematic{}
	if len(lines) == 0 {
		return result, ErrEmptySchematic
	}
	result.MaxX = len(lines[0]) - 1
	result.MaxY = len(lines) - 1
	return result, nil
}
