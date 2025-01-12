package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetring
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetring:   "retring",
}

func (s ServerState) String() string {
	return stateName[s]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetring:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
