package glean

import (
	"fmt"
	"math/rand"
	"time"
)

type Serf struct {
	goingToStart []string
	done         []string
}

func NewSerf(start []string, done []string) *Serf {
	return &Serf{goingToStart: start, done: done}
}

func (g *Serf) ExpressHowAmusedYouAreToStart() {
	rand.Seed(time.Now().Unix())

	fmt.Println("\n" + g.goingToStart[rand.Intn(len(g.goingToStart))] + " 😒\n")
}

func (g *Serf) ExpressHowHappyYouAreWithResults() {
	rand.Seed(time.Now().Unix())

	fmt.Println("\n😒" + g.done[rand.Intn(len(g.done))] + "\n")
}
