package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

const (
	A string = "A"
	B string = "B"
	C string = "C"
	D string = "D"
	E string = "E"
	F string = "F"
	G string = "G"
	H string = "H"
)

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	retrievedFile, ok := cb[file]
	if !ok {
		return 0
	} else {
		totalCount := 0
		for _, val := range retrievedFile {
			if val {
				totalCount++
			}
		}
		return totalCount
	}
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	} else {
		totalCount := 0
		for _, file := range cb {
			if file[rank-1] {
				totalCount++
			}
		}
		return totalCount
	}
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	fileNames := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	totalCount := 0
	for range fileNames {
		for i := 0; i < 8; i++ {
			totalCount += 1
		}
	}
	return totalCount
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	fileNames := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	totalCount := 0
	for _, fileName := range fileNames {
		totalCount += CountInFile(cb, fileName)
	}
	return totalCount
}
