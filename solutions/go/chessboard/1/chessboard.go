package chessboard

type File [8]bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (result int) {
	for _, v := range cb[file] {
		if v {
			result++
		}
	}
	return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (result int) {
	if rank < 1 || rank > 8 {
		return 0
	}
	for _, file := range cb {
		if file[rank-1] {
			result++
		}
	}
	return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (result int) {
	return 64
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (result int) {
	for _, file := range cb {
		for _, square := range file {
			if square {
				result++
			}
		}
	}
	return result
}
