package main

type change struct {
	i, j, val int
}

func gameOfLife(board [][]int) {
	changes := make([]change, 0, len(board)*len(board[0]))
	for i, r := range board {
		for j := range r {
			changes = append(changes, checkNeighbors(board, i, j))
		}
	}
	for _, ch := range changes {
		board[ch.i][ch.j] = ch.val
	}
}

func checkNeighbors(board [][]int, i, j int) change {
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
			if board[ii][jj] == 1 {
				liveNeighbors++
			}
		}
	}
	if liveNeighbors < 2 || liveNeighbors > 3 {
		return change{i, j, 0}
	}
	if liveNeighbors == 3 {
		return change{i, j, 1}
	}
	return change{i, j, board[i][j]}
}
