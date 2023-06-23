package main

import (
	"fmt"
	"time"
)

func main()  {
	sm := NewStateMachine(RedLight{})
	for {
		sm.Transition()
	}
}

// Define the methods that each state must implement.
type State interface {
	// when a state is entered
	Enter()
	// called when a state is exited
	Exit()
	// update the state machine
	Update(l *StateMachine)
}

// Represent the state machine to maintain the states and transitions between them.
type StateMachine struct {
	currentState State
	states       map[string]State
}

// Factory creates a new StateMachine
func NewStateMachine(initialState State) *StateMachine {
	sm := &StateMachine{
		currentState: initialState,
		states:       make(map[string]State),
	}

	sm.currentState.Enter()
	return sm
}

// Set initial state of the state machine
func (sm *StateMachine) setState(s State) {
	sm.currentState = s
	sm.currentState.Enter()
}

// Update the state machine
func (sm *StateMachine) Transition() {
	sm.currentState.Update(sm)
}

type RedLight struct {}

func (g RedLight) Enter() {
	fmt.Println("Red light is on. Stop driving.")
	time.Sleep(time.Second * 5)
}

func (g RedLight) Exit() {}

func (g RedLight) Update(l *StateMachine) {
	l.setState(&Greenligth{})
}

type Greenligth struct {}

func (g Greenligth) Enter() {
	fmt.Println("Green light is on. You can drive.")
	time.Sleep(time.Second * 5)
}

func (g Greenligth) Exit() {}

func (g Greenligth) Update(l *StateMachine) {
	l.setState(&YellowLight{})
}

type YellowLight struct {}

func (g YellowLight) Enter() {
	fmt.Println("Yellow light is on. Prepare to stop.")
	time.Sleep(time.Second * 5)
}

func (g YellowLight) Exit() {}

func (g YellowLight) Update(l *StateMachine) {
	l.setState(&RedLight{})
}

