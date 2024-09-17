package main

import "fmt"

// same as typedef
type ServerState int

// const (...) declares a set of constant values
const (
	// iota generates succerssive const vals automatically (here 0, 1, 2, 3)
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{

	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying", // note that we still need to use a comma on the last entry
}

// stringer func to modify output
// stringers cannot take bare ints
func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unkown state: %s", s))
	}
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(StateIdle)
	fmt.Println(ns2)
}
