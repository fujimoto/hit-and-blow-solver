package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N_TARGET = 6
const N_CHOICE = 4
const IS_DUPLICATE_CHOICE = true
const N_MAX_TRY = 1024

func main() {
	answer := generateAnswer()
	fmt.Printf("answer: %v\n", answer)

	solve(answer)
}

func generateAnswer() [N_CHOICE]int {
	answer := [N_CHOICE]int{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N_CHOICE; i++ {
		answer[i] = rand.Intn(N_TARGET)

		if IS_DUPLICATE_CHOICE == false && i > 0 {
			duplicate := false
			for j := 0; j < i; j++ {
				if answer[i] == answer[j] {
					duplicate = true
					break
				}
			}

			if duplicate {
				i--
			}
		}
	}

	return answer
}

func solve(answer [N_CHOICE]int) {
	solved := false

	tracks := [][N_CHOICE]int{}

	try := 0
	for try = 0; try < N_MAX_TRY; try++ {
		input := infer(tracks)
		hit, blow := judge(input, answer)

		fmt.Printf("try[%d]: input=%v, hit=%d, blow=%d\n", input, hit, blow)

		if hit == N_CHOICE {
			solved = true
			break
		}
	}

	fmt.Printf("solve: solved=%v, try=%v", solved, try)
}

func judge(input [N_CHOICE]int, answer [N_CHOICE]int) (int, int) {
}
