package bowling

import (
	"testing"
	"fmt"
)

func rollMany(n, pins int, g *Game) {
	for i := 0; i < n; i++ {
		g.Roll(pins)
	}
}

func assertScore(n int, g *Game, t *testing.T) {
	score := g.Score()
	if score != n {
		t.Fail()
		fmt.Printf("Expected score %d and got %d", n, score)
	}
}

func TestZero(t *testing.T) {
	g := NewGame()
	rollMany(20, 0, g)
	assertScore(0, g, t)
}

func TestAllOnes(t *testing.T) {
	g := NewGame()
	rollMany(20, 1, g)
	assertScore(20, g, t)
}

func TestOneSpare(t *testing.T) {
	g := NewGame()
	g.Roll(5)
	g.Roll(5) // spare
	g.Roll(3)
	rollMany(17, 0, g)
	assertScore(16, g, t)
}

func TestOneStrike(t *testing.T) {
	g := NewGame()
	g.Roll(10) //strike
	g.Roll(3)
	g.Roll(4)
	rollMany(16, 0, g)
	assertScore(24, g, t)
}

func TestPerfectGame(t *testing.T) {
	g := NewGame()
	rollMany(12, 10, g)
	assertScore(300, g, t)
}
