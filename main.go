package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

const diceNotationPattern = `^(\d+)d(\d+)$`

var diceRegex = regexp.MustCompile(diceNotationPattern)

var (
	ErrInvalidDiceNotation = errors.New("invalid dice notation")
	ErrInvalidRollNumber   = errors.New("invalid number of dice rolls")
	ErrInvalidDiceSides    = errors.New("invalid dice sides")
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage:\t\t%s <notation...>\n", os.Args[0])
		fmt.Printf("Example:\t%s 1d6 2d20\n", os.Args[0])
		os.Exit(0)
	}

	for _, arg := range os.Args[1:] {
		rollNumber, diceSides, err := parseDiceNotation(arg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		results := rollDice(rollNumber, diceSides)
		fmt.Printf("%s: %v: %d\n", arg, results, sum(results))
	}
	return
}

func parseDiceNotation(notation string) (rollNumber int, diceSides int, err error) {
	matches := diceRegex.FindStringSubmatch(notation)
	if matches == nil {
		return 0, 0, fmt.Errorf("%s: %w", notation, ErrInvalidDiceNotation)
	}

	rollNumber, err = strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", notation, ErrInvalidRollNumber)
	}

	diceSides, err = strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", notation, ErrInvalidDiceSides)
	}

	return rollNumber, diceSides, nil
}

func rollDice(rollNumber int, diceSides int) (results []int) {
	results = make([]int, rollNumber)
	for i := 0; i < rollNumber; i++ {
		results[i] = rand.Intn(diceSides) + 1
	}
	return results
}

func sum(values []int) (total int) {
	total = 0
	for _, value := range values {
		total += value
	}
	return total
}
