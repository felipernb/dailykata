package bowling

type Game struct {
	rolls []int
	currentRoll int
}

func NewGame() *Game {
	return &Game{make([]int, 21), 0}
}

func (g *Game) Roll(p int) {
	g.rolls[g.currentRoll] = p
	g.currentRoll++
}

func (g *Game) Score() int {
	score := 0
	frame := 0
	for i := 0; i < 10; i++ {
		if g.isStrike(frame) {
			score += 10 + g.strikeBonus(frame)
			frame++
		} else if g.isSpare(frame) {
			score += 10 + g.spareBonus(frame)
			frame += 2
		} else {
			score += g.sumOfRolls(frame)
			frame += 2
		}
	}
	return score
}

func (g *Game) isStrike(frame int) bool {
	return g.rolls[frame] == 10
}

func (g *Game) isSpare(frame int) bool {
	return g.sumOfRolls(frame) == 10
}

func (g *Game) strikeBonus(frame int) int {
	return g.sumOfRolls(frame+1)
}

func (g *Game) spareBonus(frame int) int {
	return g.rolls[frame+2]
}

func (g *Game) sumOfRolls(frame int) int {
	return g.rolls[frame] + g.rolls[frame+1]
}
