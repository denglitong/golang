// Go supports first class functions, higher-order functions,
// user-definced function types, function literals, closures,
// and multiple return values.
// The rich feature set supports a functional programming style
// in a strongly typed language.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	win            = 100 // The wining score in a game of Pig
	gamesPerSeries = 10  // The number of games per series to simulate
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// score A score includes scores accumulated in previous turns
// for each player as well as the points scored by the current player in this turn.
type score struct {
	player, opponent, thisTurn int
}

// action An action transitions stochastically to a resulting score.
type action func(current score) (result score, turnIsOver bool)

// roll returns the (result, turnIsOver) outcome of a simulating a die roll.
// If the roll value is 1, then thisTurn score is abandoned, and the player's
// roles swap. Otherwise, the roll value is added to thisTurn.
func roll(s score) (score, bool) {
	outcome := random.Intn(6) + 1 // A random int in [1,6]
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}
	return score{s.player, s.opponent, outcome + s.thisTurn}, false
}

// stay returns the (result, turnIsOver) outcome of staying.
// thisTurn score is added to the player's score, and the player's roles swap.
func stay(s score) (score, bool) {
	return score{s.opponent, s.player + s.thisTurn, 0}, true
}

// strategy A strategy choose an action for any given socre.
type strategy func(score) action

// stayAtK returns a strategy that rolls until thisTurn is at least k, then stays.
func stayAtK(k int) strategy {
	return func(s score) action {
		if s.thisTurn >= k {
			return stay
		}
		return roll
	}
}

// play simulates a Pig game and returns the winner (0 or 1)
func play(strategy0, strategy1 strategy) int {
	strategies := []strategy{strategy0, strategy1}
	s, turnIsOver := score{}, false
	currentPlayer := random.Intn(2) // Randomly decide who plays first
	for s.player+s.thisTurn < win {
		action := strategies[currentPlayer](s)
		s, turnIsOver = action(s)
		if turnIsOver {
			// transit the player to start in next turn
			currentPlayer = (currentPlayer + 1) % 2
		}
	}
	return currentPlayer
}

// roundRobin simulates a series of games between every pair of strategies.
func roundRobin(strategies []strategy) ([]int, int) {
	wins := make([]int, len(strategies))
	for i := 0; i < len(strategies); i++ {
		for j := i + 1; j < len(strategies); j++ {
			for k := 0; k < gamesPerSeries; k++ {
				winner := play(strategies[i], strategies[j])
				if winner == 0 {
					wins[i]++
				} else {
					wins[j]++
				}
			}
		}
	}
	gamesPerStrategy := gamesPerSeries * (len(strategies) - 1) // no self play
	return wins, gamesPerStrategy
}

// ratioString takes a list of integer values and returns a string that
// lists each value and its percentage of the sum of all values.
// e.g., ratios(1, 2, 3) = "1/6 (17.17%), 2/6 (33.3%), 3/6 (50.9%)"
func ratioString(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	s := ""
	for _, val := range vals {
		if s != "" {
			s += ", "
		}
		pct := 100 * float64(val) / float64(total)
		s += fmt.Sprintf("%d/%d (%0.1f%%)", val, total, pct)
	}
	return s
}

func main() {
	strategies := make([]strategy, win)
	for k := range strategies {
		strategies[k] = stayAtK(k + 1)
	}

	wins, games := roundRobin(strategies)

	for k := range strategies {
		fmt.Printf("Staying at k =% 4d: wins, losses %s\n",
			k+1, ratioString(wins[k], games-wins[k]))
	}
}
