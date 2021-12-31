package chessboard

import (
	"strings"
)

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [8]bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	//panic("Please implement CountInRank()")
	if strings.ToUpper(rank) >= "A" && strings.ToUpper(rank) <= "H" {
		var sum int
		for _, v := range cb[strings.ToUpper(rank)] {
			if v {
				sum++
			}
		}
		return sum
	}
	return 0
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	//panic("Please implement CountInFile()")
	if file >= 1 && file <= 8 {
		var sum int
		for _, v := range cb {
			if v[file-1] {
				sum++
			}
		}
		return sum
	}
	return 0
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	//panic("Please implement CountAll()")
	var t int
	for range cb {
		t += 8
	}
	return t
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	//panic("Please implement CountOccupied()")
	var t int
	for k, _ := range cb {
		t += CountInRank(cb, k)
	}
	return t
}
