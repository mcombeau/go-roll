package main

import (
	"errors"
	"testing"
)

func TestParseDiceNotation(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		wantRollNumber int
		wantDiceSides  int
		wantError      error
	}{
		{
			name:           "Test valid dice notation",
			input:          "2d20",
			wantRollNumber: 2,
			wantDiceSides:  20,
			wantError:      nil,
		},
		{
			name:           "Test empty dice notation",
			input:          "",
			wantRollNumber: 0,
			wantDiceSides:  0,
			wantError:      ErrInvalidDiceNotation,
		},
		{
			name:           "Test invalid dice notation: 2d",
			input:          "2d",
			wantRollNumber: 0,
			wantDiceSides:  0,
			wantError:      ErrInvalidDiceNotation,
		},
		{
			name:           "Test invalid number of dice rolls",
			input:          "100000000000000000000000000000000000000000d20",
			wantRollNumber: 0,
			wantDiceSides:  0,
			wantError:      ErrInvalidRollNumber,
		},
		{
			name:           "Test invalid dice sides",
			input:          "1d100000000000000000000000000000000000000000",
			wantRollNumber: 0,
			wantDiceSides:  0,
			wantError:      ErrInvalidDiceSides,
		},
		{
			name:           "Test double dice notation",
			input:          "1d23d4",
			wantRollNumber: 0,
			wantDiceSides:  0,
			wantError:      ErrInvalidDiceNotation,
		},
		{
			name:           "Test invalid dice notation: 1dd2",
			input:          "2dd2",
			wantRollNumber: 0,
			wantDiceSides:  0,
			wantError:      ErrInvalidDiceNotation,
		},
		{
			name:           "Test invalid dice notation: d",
			input:          "d",
			wantRollNumber: 0,
			wantDiceSides:  0,
			wantError:      ErrInvalidDiceNotation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRollNumber, gotDiceSides, gotError := parseDiceNotation(tt.input)

			if tt.wantError != nil {
				if gotError == nil || !errors.Is(gotError, tt.wantError) {
					t.Fatalf("parseDiceNotation error got = %v, want error = %v, input = %s\n", gotError.Error(), tt.wantError.Error(), tt.input)
				}
			}

			if gotRollNumber != tt.wantRollNumber {
				t.Errorf("parseDiceNotation rollNumber got = %d, want = %d, input = %s\n", gotRollNumber, tt.wantRollNumber, tt.input)
			}

			if gotDiceSides != tt.wantDiceSides {
				t.Errorf("parseDiceNotation diceSides got = %d, want = %d, input = %s\n", gotDiceSides, tt.wantDiceSides, tt.input)
			}
		})
	}
}
