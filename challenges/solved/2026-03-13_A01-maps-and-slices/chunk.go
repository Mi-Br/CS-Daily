package main

// ### Exercise 3: Chunk

// ```go
// // Chunk splits a slice into sub-slices of the given size.
// // The last chunk may be smaller if len(s) is not divisible by n.
// //
// // Example: Chunk([]int{1,2,3,4,5}, 2) → [[1 2] [3 4] [5]]
// // Example: Chunk([]int{1,2,3}, 3)     → [[1 2 3]]
// // Example: Chunk([]int{}, 2)          → []
// func Chunk(s []int, n int) [][]int
// ```

// **Rules:**
// - Each chunk must be an independent copy — modifying a chunk must NOT affect the original slice
// - Handle edge cases: empty slice, n larger than len(s), n = 1
// - No aliasing. This is the point of this exercise.

func Chunk(inp []int, size int) [][]int {

	outp := [][]int{}
	if size == 0 || len(inp) == 0 {
		return outp
	}
	for i := 0; i < len(inp); i += size {
		end := i + size
		if end > len(inp) {
			end = len(inp)
		}
		segment := inp[i:end]
		chunkCopy := make([]int, len(segment))
		copy(chunkCopy, segment)
		outp = append(outp, chunkCopy)
	}

	return outp
}
