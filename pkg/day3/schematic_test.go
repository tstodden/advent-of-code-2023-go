package day3_test

import (
	"aoc2023/pkg/day3"
	"testing"
)

func TestNewSchematic_SingleLine(t *testing.T) {
	input := []string{"440.......#34"}

	got := day3.NewSchematic(input)

	if got.MaxX != 12 || got.MaxY != 0 {
		t.Errorf("invalid maximum dimensions: x = %d and y = %d", got.MaxX, got.MaxY)
	}
	if length := len(got.References); length != 2 {
		t.Errorf("invalid number of references: len = %d", length)
		return
	}
	if ref := got.References[0]; ref.IsPart != false || ref.Value != 440 {
		t.Errorf("invalid first reference: %#v", ref)
	}
	if ref := got.References[1]; ref.IsPart != true || ref.Value != 34 {
		t.Errorf("invalid first reference: %#v", ref)
	}
}
