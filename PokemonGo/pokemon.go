package main

import (
	"PokemonGo/Models"
	"bufio"
	"fmt"
	"os"
)

const comma = ","

func GottaCatchEmAll(moves string) (int, error) {

	//Use a map to check if a given coordinate has been visited
	//Map is of nature: key -> "x,y", val -> true/false
	//In Java you could probably use an HashSet for O(1) lookup and unique positions, but Go does not have it :(
	visitedSpots := make(map[string]bool)

	var (
		pos          = Models.Position{}       // Initial position with x:0 and y:0
		count        = 1                       // Number of caught pokemon, start at 1
		currPosition = fmt.Sprint(0, comma, 0) // Current Position of Ash, start at "0,0"

	)

	//Mark initial position as visited
	visitedSpots[currPosition] = true

	//Iterate over moves
	for _, move := range moves {

		err := pos.EvaluateMoveBranching(move)

		//Catch err if illegal moves was made
		if err != nil {
			return 0, err
		}

		//Update current position
		currPosition = fmt.Sprint(pos.X, comma, pos.Y)

		//Check if map contains current position
		//If not -> add to map and catch pokemon
		//If it has, do nothing and move on to next iteration
		if _, contains := visitedSpots[currPosition]; !contains {
			visitedSpots[currPosition] = true
			count++
		}
	}

	return count, nil
}

func main() {

	//Read moves from Stdin until a lineBreak
	in := bufio.NewScanner(os.Stdin)

	in.Scan()

	moves := in.Text()
	//Proceed to the logic
	count, err := GottaCatchEmAll(moves)

	//Check for error in the logic function, namely invalid moves out of the {N,S,E,O}
	checkError(err)

	//Print result
	fmt.Println(count)
}

//Check for unrecoverable errors, throwing a panic (program ends) if an error exists
func checkError(err error) {

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
