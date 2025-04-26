package main

const (
	reviveMe = iota - 1
	death
	alive
	killMe
)

// https://leetcode.com/problems/game-of-life
// O(n^2) time
// O(1) space
func gameOfLife(board [][]int) {
	for i, r := range board {
		for j := range r {
			checkNeighbors(board, i, j)
		}
	}
	for i, r := range board {
		for j, val := range r {
			if val == reviveMe {
				board[i][j] = 1
			} else if val == killMe {
				board[i][j] = 0
			}
		}
	}
}

func checkNeighbors(board [][]int, i, j int) {
	m := len(board)
	n := len(board[0])
	liveNeighbors := 0
	for ii := i - 1; ii <= i+1; ii++ {
		if ii < 0 || ii >= m {
			continue
		}
		for jj := j - 1; jj <= j+1; jj++ {
			if jj < 0 || jj >= n || (ii == i && jj == j) {
				continue
			}
			if board[ii][jj] >= alive {
				liveNeighbors++
			}
		}
	}
	if board[i][j] == alive && (liveNeighbors < 2 || liveNeighbors > 3) {
		board[i][j] = killMe
	}
	if board[i][j] == death && liveNeighbors == 3 {
		board[i][j] = reviveMe
	}
}
