package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearScreen() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func drawBoard(b [9]int) {
	for i, v := range b {
		if v == 0 {
			fmt.Print(i + 1)
		} else if v == 1 {
			fmt.Print("O")
		} else if v == 2 {
			fmt.Print("X")
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Println()
		} else {
			fmt.Print(" | ")
		}
	}
}

func AIRandomNextMove(b [9]int) int {
	choices := make([]int, 0, 9)
	for i, v := range b {
		if v == 0 {
			choices = append(choices, i)
		}
	}

	return choices[rand.Intn(len(choices))]
}

func main() {
	ai := flag.Bool("ai", false, "to play vs ai")
	flag.Parse()

	rand.Seed(time.Now().Unix())

	gameOver := false
	winner := 0

	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	turn := 1
	var status string

	for {
		clearScreen()
		drawBoard(board)
		if gameOver {
			if winner == 1 {
				fmt.Println("O - winner")
			} else if winner == 2 {
				fmt.Println("X - winner")
			} else {
				fmt.Println("DRAW")
			}
			return
		}
		player := turn%2 + 1
		if player == 1 {
			fmt.Println("O turn")
		} else {
			fmt.Println("X turn")
		}

		fmt.Println("Select a move")

		if status != "" {
			fmt.Println(status)
		}

		if player == 2 || !*ai {
			var moveIndex string
			fmt.Scan(&moveIndex)

			if moveIndex == "q" {
				fmt.Println("Exiting")
				os.Exit(0)
			}

			switch moveIndex {
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":
				idx, _ := strconv.Atoi(string(moveIndex))
				fmt.Println(idx)
				idx--
				if board[idx] != 0 {
					status = fmt.Sprintf("%v is already occupied", idx+1)
					continue
				} else {
					board[idx] = player
				}
				status = ""
				break
			default:
				status = "1-9 or q to quit"
				continue
			}
		} else {
			board[AIRandomNextMove(board)] = player
		}

		turn++

		winCombinations := [8][3]int{
			{0, 1, 2},
			{3, 4, 5},
			{6, 7, 8},
			{0, 3, 6},
			{1, 4, 7},
			{2, 5, 8},
			{0, 4, 8},
			{6, 4, 2},
		}

		for _, c := range winCombinations {
			if board[c[0]] != 0 && board[c[0]] == board[c[1]] && board[c[0]] == board[c[2]] {
				gameOver = true
				winner = board[c[0]]
				break
			}
		}

		if turn == 10 {
			gameOver = true
		}
	}
}
