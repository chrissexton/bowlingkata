package bowling

import "fmt"

type Game struct {
	initialScore int
	frames       []*frame
}

type frame struct {
	rolls []int
}

func mkFrame() *frame {
	return &frame{
		rolls: []int{},
	}
}

// score returns score, is a strike, and is a spare
func (f frame) score() (int, bool, bool) {
	if len(f.rolls) == 0 {
		return 0, false, false
	}
	if len(f.rolls) == 1 {
		return f.rolls[0], f.rolls[0] == 10, false
	}
	score := 0
	for _, roll := range f.rolls {
		score += roll
	}
	return score, false, score == 10
}

func New(startingScore int) *Game {
	return &Game{
		initialScore: startingScore,
		frames: []*frame{
			mkFrame(),
		},
	}
}

func (g *Game) Roll(pins int) {
	lastFrame := g.lastFrame()
	_, isStrike, _ := lastFrame.score()
	if len(lastFrame.rolls) >= 2 || isStrike {
		g.frames = append(g.frames, mkFrame())
		lastFrame = g.lastFrame()
	}
	lastFrame.rolls = append(lastFrame.rolls, pins)
}

func (g *Game) Score() int {
	score := 0
	for i, frame := range g.frames {
		// If we are at the end of the game
		if i == 10 {
			break
		}
		frameScore, isStrike, isSpare := frame.score()
		if isSpare {
			score += g.frames[i+1].rolls[0]
		}
		if isStrike && len(g.frames) > i+1 {
			score += g.frames[i+1].rolls[0]
			// If the next roll is a strike, there is only one roll in the current next frame
			if _, nextStrike, _ := g.frames[i+1].score(); nextStrike {
				score += g.frames[i+2].rolls[0]
			} else {
				score += g.frames[i+1].rolls[1]
			}
		}
		score += frameScore
	}
	return score
}

func (g *Game) lastFrame() *frame {
	return g.frames[len(g.frames)-1]
}

func (g *Game) String() string {
	out := ""
	for i, frame := range g.frames {
		out += fmt.Sprintf("| %d: ", i)
		for _, roll := range frame.rolls {
			out += fmt.Sprintf("%d ", roll)
		}
	}
	return out
}
