package intcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Machine is the intcode-machine
type Machine struct {
	State    []int
	Position int
}

// New creates a new intcode-machine with the specified
// program supplied as an int-slice
func New(state []int, pos int) *Machine {
	return &Machine{
		State:    state,
		Position: pos,
	}
}

func (m *Machine) Step() {
	instr := m.State[m.Position]
	switch instr % 100 {
	case 1:
		m.Add()
	case 2:
		m.Mult()
	case 3:
		m.Input()
	case 4:
		m.Output()
	case 99:
		m.Halt()
	}

}

func (m *Machine) Add() {
	param1 := 0
	param2 := 0
	instr := m.State[m.Position]
	instr = instr - 1
	if instr%1000 == 100 {
		param1 = m.State[m.Position+1]
		instr = instr - 100
	} else {
		param1 = m.State[m.State[m.Position+1]]
	}
	if instr%10000 == 1000 {
		param2 = m.State[m.Position+2]
	} else {
		param2 = m.State[m.State[m.Position+2]]
	}
	m.State[m.State[m.Position+3]] = param1 + param2
	m.Position = m.Position + 4
}

func (m *Machine) Mult() {
	param1 := 0
	param2 := 0
	instr := m.State[m.Position]
	instr = instr - 1
	if instr%1000 == 100 {
		param1 = m.State[m.Position+1]
		instr = instr - 100
	} else {
		param1 = m.State[m.State[m.Position+1]]
	}
	if instr%10000 == 1000 {
		param2 = m.State[m.Position+2]
	} else {
		param2 = m.State[m.State[m.Position+2]]
	}
	m.State[m.State[m.Position+3]] = param1 * param2
	m.Position = m.Position + 4
}

func (m *Machine) Input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give input: ")
	value, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	m.State[m.State[m.Position+1]] = intValue

}

func (m *Machine) Output() {
	outputVal := m.State[m.State[m.Position+1]]
	fmt.Println(outputVal)
	m.Position = m.Position + 2
}

func (m *Machine) Halt() {
	fmt.Println(m.State)

}
