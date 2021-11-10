package main

import "fmt"

//Moves allowed
const (
	north = 'N'
	south = 'S'
	east  = 'E'
	west  = 'O'
)

const unknownMove = "move '%c' not supported. Supported moves: [N,S,E,O]"

//Position struct containing X and Y position of Ash
type Position struct {
	X int
	Y int
}

//Action methods which move Ash and update his location.
//Y axis goes North and South
//X axis goes West and East
func (pos *Position) goSouth() {
	pos.Y -= 1
}
func (pos *Position) goNorth() {
	pos.Y += 1
}
func (pos *Position) goWest() {
	pos.X -= 1
}
func (pos *Position) goEast() {
	pos.X += 1
}

//Global Map which maps a certain move to its respective action
//E.g -> If 'N' then apply goNorth
var actions = make(map[int32]func(pos *Position))

//Initialize moves and actions
func init() {
	actions[north] = (*Position).goNorth
	actions[south] = (*Position).goSouth
	actions[east] = (*Position).goEast
	actions[west] = (*Position).goWest
}

//Evaluates a move
//If move is present in actions map, then perform its corresponding action
//If move not present in actions map, then return an error for illegal move
func (pos *Position) EvaluateMoveMapped(move int32) error {

	if action, contains := actions[move]; contains {

		action(pos)

	} else {
		return fmt.Errorf(unknownMove, move) // If unknown move, throw error
	}

	return nil
}

func (pos *Position) EvaluateMoveBranching(move int32) error {

	switch move {
	case north:
		pos.goNorth()
	case south:
		pos.goSouth()
	case east:
		pos.goEast()
	case west:
		pos.goWest()
	default:
		return fmt.Errorf(unknownMove, move) // If unknown move, throw error
	}

	return nil
}
