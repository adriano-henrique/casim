package games

import (
	"fmt"
	"math/rand"
	"strconv"
)

func SimulateRoulleteGame(numberOfGames int) {
	for i := 0; i < numberOfGames; i++ {
		randomNumber := rand.Intn(38) + 1
		resultValue := getResultingValueString(randomNumber)
		fmt.Printf("Game %d: %s \n", i, resultValue)
	}
}

func getResultingValueString(randomNumber int) string {
	if randomNumber == 37 {
		return "0"
	}

	if randomNumber == 38 {
		return "00"
	}

	return strconv.Itoa(randomNumber)
}
