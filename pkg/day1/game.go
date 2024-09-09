package day1

import (
	"errors"
	"strconv"
	"strings"
)

var ErrGameFormat = errors.New("color must have exactly 2 parts")

type Game struct {
	Id      int
	Samples []Sample
}

func NewGame(s string) (Game, error) {
	var result Game

	parts := strings.Split(s, ": ")
	id, err := getGameId(parts[0])
	if err != nil {
		return result, err
	}
	result.Id = id

	samples := strings.Split(parts[1], "; ")
	for _, v := range samples {
		sample, err := NewSample(v)
		if err != nil {
			return result, err
		}
		result.Samples = append(result.Samples, sample)
	}

	return result, nil
}

func getGameId(s string) (int, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return -1, ErrGameFormat
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, err
	}
	return id, nil
}
