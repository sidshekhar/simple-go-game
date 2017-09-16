
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Welcome         string
	Energy          int
	CurrentLocation string
}

func (g *Game) Play() {
	fmt.Println(g.Welcome)
	for {
		fmt.Println(locationMap[g.CurrentLocation].Description)
		g.ProcessEvents(locationMap[g.CurrentLocation].Events)
		if g.Energy <= 0 {
			fmt.Println("You are wiped out!! The manager calls for you to be substituted!")
			return
		}
		if g.Energy > 100 {
			fmt.Println("You won the game! The trophy is yours... Congratulations!")
			return
		}
		fmt.Printf("Energy: %d\n", g.Energy)
		fmt.Println("You can go to these places:")
		for index, loc := range locationMap[g.CurrentLocation].Transitions {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(locationMap[g.CurrentLocation].Transitions) {
			fmt.Printf("%s%d%s\n", "Where do you want to go (0 - to quit), [1...", len(locationMap[g.CurrentLocation].Transitions), "]: ")
			fmt.Scan(&i)
		}
		newLoc := i - 1
		g.CurrentLocation = locationMap[g.CurrentLocation].Transitions[newLoc]

	}
}

func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.Energy += evts[evtName].ProcessEvent()
	}
}

type Event struct {
	Type        string
	Chance      int
	Description string
	Energy      int
	Evt         string
}

func (e *Event) ProcessEvent() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if e.Chance >= r1.Intn(100) {
		hp := e.Energy
		if e.Type == "Combat" {
			fmt.Println("Tackle!")
		}
		fmt.Printf("\t%s\n", e.Description)
		if e.Evt != "" {
			hp = hp + evts[e.Evt].ProcessEvent()
		}
		return hp
	}
	return 0
}

type Location struct {
	Description string
	Transitions []string
	Events      []string
}

var evts = map[string]*Event{
	"standingTackle": {Type: "Combat", Chance: 40, Description: "A tough defender stands in your way with a firm tackle.", Energy: -35, Evt: "teammateRecovery"},
	"teammateRecovery": {Type: "Story", Chance: 10, Description: "Your teammate rushes to your rescue and distracts the defenders away from you.", Energy: +30, Evt: ""},
	"slideTackle": {Type: "Combat", Chance: 40, Description: "A frustrated defender lunges at you with a slide tackle!", Energy: -60, Evt: ""},
	"recoverEnergy": {Type: "Story", Chance: 100, Description: "While playing defense, you have a chance to catch your breath a little.", Energy: +10, Evt: ""},
	"scoreGoal": {Type: "Story", Chance: 80, Description: "Being the superstar you are, you score a goal!", Energy: +20, Evt: ""},

}

var locationMap = map[string]*Location{
	"Midfield":      {"You are the best player on your team standing in Midfield with the ball", []string{"Left Wing", "Right Wing"}, []string{"standingTackle"}},
	"Right Wing":  {"You are now running up the right wing.", []string{"Midfield", "Penalty Box"}, []string{}},
	"Left Wing":  {"You are now toying with the ball in the left flank.", []string{"Midfield", "Penalty Box", "Defensive"}, []string{"slideTackle"}},
	"Penalty Box": {"You are in prime scoring position!", []string{"Right Wing", "Left Wing"}, []string{"scoreGoal"}},
	"Defensive": {"You are now in a safe defensive area.", []string{"Midfield", "Right Wing", "Left Wing"}, []string{"recoverEnergy"}},
}

func main() {
	g := &Game{Energy: 100, Welcome: "Welcome to the Champions League Final\n", CurrentLocation: "Midfield"}
	g.Play()
}