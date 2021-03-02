package bowling

import (
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
	messages "github.com/cucumber/messages-go/v10"
)

var g *Game

func thePlayerRollsATimes(pins, times int) error {
	for i := 0; i < times; i++ {
		g.Roll(pins)
	}
	return nil
}

func theScoreShouldBe(expectedScore int) error {
	// fmt.Println(g.String())
	actualScore := g.Score()
	if actualScore != expectedScore {
		return fmt.Errorf("the score %d did not match the expected score of %d", actualScore, expectedScore)
	}
	return nil
}

func theScoreStartsAt(initialScore int) error {
	g = New(initialScore)
	return nil
}

func thePlayerRollsA(tbl *messages.PickleStepArgument_PickleTable) error {
	for _, row := range tbl.GetRows() {
		roll, err := strconv.Atoi(row.Cells[0].Value)
		if err != nil {
			return err
		}
		g.Roll(roll)
	}
	return nil
}

func wePrintOutTheGame() error {
	fmt.Println(g.String())
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	g = nil
	ctx.Step(`^the player rolls a (\d+) (\d+) times$`, thePlayerRollsATimes)
	ctx.Step(`^the score should be (\d+)$`, theScoreShouldBe)
	ctx.Step(`^the score starts at (\d+)$`, theScoreStartsAt)
	ctx.Step(`^the player rolls a$`, thePlayerRollsA)
	ctx.Step(`^we print out the game$`, wePrintOutTheGame)
}
