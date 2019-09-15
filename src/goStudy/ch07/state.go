package main

import (
	"errors"
	"fmt"
	"reflect"
)

type State interface {
	Name() string
	EnableSameTransit() bool
	OnBegin()
	OnEnd()
	CanTransitTo(name string) bool
}

func StateName(s State) string {
	if s == nil {
		return "none"
	}
	return reflect.TypeOf(s).Elem().Name()
}

type StateInfo struct {
	name string
}

func (s *StateInfo) Name() string {
	return s.name
}
func (s *StateInfo) setName(name string) {
	s.name = name
}

func (s *StateInfo) EnableSameTransit() bool {
	return false
}

func (s *StateInfo) OnBegin() {

}

func (s *StateInfo) OnEnd() {

}

func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}

//状态管理
type StateManager struct {
	stateByName map[string]State
	OnChange    func(from, to State)
	curr        State
}

func (sm *StateManager) Add(s State) {
	name := StateName(s)
	s.(interface {
		setName(name string)
	}).setName(name)
	if sm.Get(name) != nil {
		panic("duplicate state:" + name)
	}
	sm.stateByName[name] = s
}

func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
		OnChange:    nil,
		curr:        nil,
	}
}

var ErrStateNotFound = errors.New("state not found")
var ErrForbidStateTransit = errors.New("forbid same state transit")
var ErrCannotTransitToState = errors.New("cannot transit to state")

func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}
	return sm.curr.CanTransitTo(name)
}

func (sm *StateManager) Transit(name string) error {
	next := sm.Get(name)
	if next != nil {
		return ErrStateNotFound
	}
	pre := sm.curr
	if sm.curr != nil {
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrForbidStateTransit
		}
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
		sm.curr.OnEnd()
	}
	sm.curr = next
	sm.curr.OnBegin()
	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}
	return nil
}

type IdleState struct {
	StateInfo
}

func (i *IdleState) OnBegin() {
	fmt.Println("IdleState begin")
}

func (i *IdleState) OnEnd() {
	fmt.Println("IdleState end")
}

type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin() {
	fmt.Println("MoveState begin")
}

func (m *MoveState) EnableSameTransit() bool {
	return true
}

type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	fmt.Println("JumpState begin")
}

func (j *JumpState) CanTransitTo(name string) bool {
	return name != "MoveState"
}

func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED!%s-->%s\n\n", sm.curr, target, err.Error())
	}
}

func main() {
	sm := NewStateManager()
	sm.OnChange = func(from, to State) {
		fmt.Printf("%s--->%s\n\n", StateName(from), StateName(to))
	}
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))
	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")
}
