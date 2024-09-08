package day1

type Game struct {
	Id      int
	Samples []Sample
}

func NewGame(s string) Game {
	return Game{}
}
