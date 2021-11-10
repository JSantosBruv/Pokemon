package main

import (
	"testing"
)

func TestGivenTestCases(t *testing.T) {

	moves := "E"
	expectedCount := 2

	count, err := GottaCatchEmAll(moves)

	checkForNormalExecution(err, count, expectedCount, moves, t)

	moves = "NESO"
	expectedCount = 4

	count, err = GottaCatchEmAll(moves)

	checkForNormalExecution(err, count, expectedCount, moves, t)

	moves = "NSNSNSNSNS"
	expectedCount = 2

	count, err = GottaCatchEmAll(moves)

	checkForNormalExecution(err, count, expectedCount, moves, t)

}

func TestCustom(t *testing.T) {

	moves := "NSSN"
	expectedCount := 3

	count, err := GottaCatchEmAll(moves)

	checkForNormalExecution(err, count, expectedCount, moves, t)

	moves = "NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN"
	expectedCount = 200

	count, err = GottaCatchEmAll(moves)

	checkForNormalExecution(err, count, expectedCount, moves, t)
}

func TestForError(t *testing.T) {

	moves := "Isto deve dar erro (hopefully)"
	expectedCount := 0

	count, err := GottaCatchEmAll(moves)

	checkForErrorHandling(err, count, expectedCount, moves, t)
}

//Assumi que se nenhuma jogada for iserida, o Ash apanha so o Pokemon da casa inicial
func TestEmpty(t *testing.T) {

	moves := ""
	expectedCount := 1

	count, err := GottaCatchEmAll(moves)

	checkForNormalExecution(err, count, expectedCount, moves, t)
}

func checkForNormalExecution(err error, count int, expectedCount int, moves string, t *testing.T) {
	if err != nil {
		t.Error(err)
	} else if count != expectedCount {
		t.Errorf("Error for move '%s': expected %d got %d", moves, expectedCount, count)
	}
}

func checkForErrorHandling(err error, count int, expectedCount int, moves string, t *testing.T) {

	if err == nil {
		t.Error("Error was not caught")
	} else if count != expectedCount {
		t.Errorf("Error for move '%s': expected %d got %d", moves, expectedCount, count)
	}
}
