package population

import (
	d "streaming-optimization/data"
	t "streaming-optimization/team"
	"math/rand"
)

// Struct for gene for genetic algorithm
type Gene struct {
	Roster  	   map[string]d.Player
	NewPlayers 	   map[string]d.Player
	Day     	   int
	Acquisitions   int
	DroppedPlayers []d.Player
	Bench 		   []d.Player
}

// Function to create a new gene
func InitGene(bt *t.BaseTeam, day int, rng *rand.Rand) *Gene {
	
	// Create a new gene
	gene := &Gene{Roster: make(map[string]d.Player), NewPlayers: make(map[string]d.Player), Day: day, Acquisitions: 0}
	
	return gene
}