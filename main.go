package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const N_TARGET = 6
const N_CHOICE = 4
const N_DUPLICATE_CHOICE = 0
const N_MAX_TRY = 1024

const N_SOLVE = 256

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < N_SOLVE; i++ {
		answer := generateAnswer()
		solve(answer)
	}
}

func generateAnswer() [N_CHOICE]int {
	answer := [N_CHOICE]int{}

	for {
		for i := 0; i < N_CHOICE; i++ {
			answer[i] = rand.Intn(N_TARGET)

		}
		if (N_CHOICE - len(unique(answer))) <= N_DUPLICATE_CHOICE {
			break
		}
	}

	return answer
}

func generateCandidates() [][N_CHOICE]int {
	c := [][N_CHOICE]int{}

	n := [N_CHOICE]int{}
	for i := 0; i < int(math.Pow(N_TARGET, float64(N_CHOICE))); i++ {
		for j := 0; j < N_CHOICE; j++ {
			n[j] = (i / int(math.Pow(N_TARGET, float64(j)))) % N_TARGET
		}

		if (N_CHOICE - len(unique(n))) <= N_DUPLICATE_CHOICE {
			c = append(c, n)
		}
	}

	return c
}

func solve(answer [N_CHOICE]int) (int, bool) {
	candidates := generateCandidates()
	inputTracks := [][N_CHOICE]int{}
	resultTracks := [][2]int{}
	var input [N_CHOICE]int

	try := 0
	for try = 0; try < N_MAX_TRY; try++ {
		input, candidates = infer(inputTracks, resultTracks, candidates)
		if candidates == nil {
			// should not happen, though
			fmt.Printf("no candidates, giving up\n")
			break
		}
		result := judge(input, answer)

		fmt.Printf("try[%d]: input=%v, hit=%d, blow=%d (answer=%v, from %d candidates)\n", try, input, result[0], result[1], answer, len(candidates))

		inputTracks = append(inputTracks, input)
		resultTracks = append(resultTracks, result)

		if result[0] == N_CHOICE {
			break
		}
	}

	return try, resultTracks[len(resultTracks)-1][0] == N_CHOICE
}

func judge(input [N_CHOICE]int, answer [N_CHOICE]int) [2]int {
	r := [2]int{0, 0}
	i := input
	w := answer

	// hit
	for j := 0; j < N_CHOICE; j++ {
		if i[j] == w[j] {
			r[0]++
			i[j] = -1
			w[j] = -1
		}
	}

	// blow
	for j := 0; j < N_CHOICE; j++ {
		if i[j] < 0 {
			continue
		}
		for k := 0; k < N_CHOICE; k++ {
			if i[j] == w[k] {
				r[1]++
				w[k] = -1
				break
			}
		}
	}

	return r
}

func infer(inputTracks [][N_CHOICE]int, resultTracks [][2]int, candidates [][N_CHOICE]int) ([N_CHOICE]int, [][N_CHOICE]int) {
	if len(inputTracks) == 0 {
		return initialInfer(), candidates
	}

	i := inputTracks[len(inputTracks)-1]
	r := resultTracks[len(resultTracks)-1]
	possibleCandidates := [][N_CHOICE]int{}

	for j := 0; j < len(candidates); j++ {
		n := candidates[j]
		if n == i {
			continue
		}

		k := possible(i, r, n)
		if k {
			possibleCandidates = append(possibleCandidates, n)
		}
	}

	if len(possibleCandidates) <= 0 {
		return [N_CHOICE]int{}, nil
	}

	return possibleCandidates[rand.Intn(len(possibleCandidates))], possibleCandidates
}

func possible(input [N_CHOICE]int, r [2]int, n [N_CHOICE]int) bool {
	j := judge(input, n)
	return j[0] == r[0] && j[1] == r[1]
}

func initialInfer() [N_CHOICE]int {
	r := [N_CHOICE]int{}

	for i := 0; i < N_CHOICE; i++ {
		if N_DUPLICATE_CHOICE > 0 {
			if i < (N_CHOICE / 2) {
				r[i] = 0
			} else {
				r[i] = 1
			}
		} else {
			r[i] = i % N_CHOICE
		}
	}

	return r
}

func unique(n [N_CHOICE]int) []int {
	m := make(map[int]bool)
	for _, v := range n {
		if !m[v] {
			m[v] = true
		}
	}

	r := make([]int, len(m))
	i := 0
	for k := range m {
		r[i] = k
		i++
	}

	return r
}
