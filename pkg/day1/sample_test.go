package day1_test

import (
	"aoc2023/pkg/day1"
	"errors"
	"strconv"
	"testing"
)

func TestNewSample_AllColors(t *testing.T) {
	input := "10 blue, 8 red, 19 green"

	sample, err := day1.NewSample(input)

	if err != nil {
		t.Errorf("error is not nil: %s", err)
		return
	}
	if sample.Blue != 10 || sample.Green != 19 || sample.Red != 8 {
		t.Errorf("Sample does not contain the expected values: %#v", sample)
	}
}

func TestNewSample_SomeColors(t *testing.T) {
	input := "8 red"

	sample, err := day1.NewSample(input)

	if err != nil {
		t.Errorf("error is not nil: %s", err)
		return
	}
	if sample.Blue != 0 || sample.Green != 0 || sample.Red != 8 {
		t.Errorf("Sample does not contain the expected values: %#v", sample)
	}
}

func TestNewSample_InvalidInteger(t *testing.T) {
	input := "eight red"

	_, err := day1.NewSample(input)

	if err == nil {
		t.Error("error is nil")
		return
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("error is incorrect: %s", err)
	}
}

func TestNewSample_InvalidFormat(t *testing.T) {
	input := "8red"

	_, err := day1.NewSample(input)

	if err == nil {
		t.Error("error is nil")
		return
	}
	if !errors.Is(err, day1.ErrColorFormat) {
		t.Errorf("error is incorrect: %s", err)
	}
}

func TestNewSample_InvalidColor(t *testing.T) {
	input := "8 yellow"

	_, err := day1.NewSample(input)

	if err == nil {
		t.Error("error is nil")
		return
	}
	if !errors.Is(err, day1.ErrInvalidColor) {
		t.Errorf("error is incorrect: %s", err)
	}
}
