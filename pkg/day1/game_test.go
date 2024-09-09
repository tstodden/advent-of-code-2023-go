package day1_test

import (
	"aoc2023/pkg/day1"
	"errors"
	"strconv"
	"testing"
)

func TestGame_IsPossible_Possible(t *testing.T) {
	game := day1.Game{Id: 1, Samples: []day1.Sample{{Blue: 1, Green: 1, Red: 1}}}

	got := game.IsPossible(day1.BagContents{Blue: 1, Green: 1, Red: 1})

	if got != true {
		t.Error("game should be possible")
	}
}

func TestGame_IsPossible_Impossible(t *testing.T) {
	game := day1.Game{Id: 1, Samples: []day1.Sample{{Blue: 3, Green: 3, Red: 3}}}

	got := game.IsPossible(day1.BagContents{Blue: 1, Green: 1, Red: 1})

	if got != false {
		t.Error("game should be impossible")
	}
}

func TestNewGame_SingleSample(t *testing.T) {
	input := "Game 72: 3 blue, 1 green, 11 red"

	game, err := day1.NewGame(input)

	if err != nil {
		t.Error("error is not nil")
		return
	}
	if game.Id != 72 {
		t.Errorf("incorrect ID: %d", game.Id)
	}
	if len(game.Samples) != 1 {
		t.Errorf("incorrect samples from game: %#v", game)
	}
}

func TestNewGame_MultipleSamples(t *testing.T) {
	input := "Game 56: 3 red, 8 blue, 11 green; 9 green, 5 red, 4 blue; 1 blue, 4 red, 4 green"

	game, err := day1.NewGame(input)

	if err != nil {
		t.Error("error is not nil")
		return
	}
	if game.Id != 56 {
		t.Errorf("incorrect ID: %d", game.Id)
	}
	if len(game.Samples) != 3 {
		t.Errorf("incorrect samples from game: %#v", game)
	}
}

func TestNewGame_InvalidId(t *testing.T) {
	input := "Game six: 3 red, 8 blue, 11 green"

	_, err := day1.NewGame(input)

	if err == nil {
		t.Error("error is nil")
		return
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("error is incorrect: %s", err)
	}
}

func TestNewGame_InvalidGamePart(t *testing.T) {
	input := "Game6: 3 red, 8 blue, 11 green"

	_, err := day1.NewGame(input)

	if err == nil {
		t.Error("error is nil")
		return
	}
	if !errors.Is(err, day1.ErrGameFormat) {
		t.Errorf("error is incorrect: %s", err)
	}
}
