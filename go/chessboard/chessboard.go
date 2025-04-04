package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	count := 0

	for _, value := range(cb[file]) {
		if value {
			count++
		}

	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	total := 0
	idx := rank -1 //array is 0 base

	for _, file := range cb {
		if idx >= 0 && idx < len(file){
			if file[idx] {
			 total++
			}
		}
	}
	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		count += len(file)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	
	for _, file := range cb {
		for _, isOccupied := range file {
			if isOccupied {
				count++
			}
		}
	}
	return count
}
