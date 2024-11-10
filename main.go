package main

import (
	"errors"
	"fmt"
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
	for _, arg := range os.Args[1:] {
		rollNumber, diceSides, err := parseDiceNotation(arg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Arg: %s, Rolls: %d, Sides: %d\n", arg, rollNumber, diceSides)
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
