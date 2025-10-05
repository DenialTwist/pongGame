package main

import (
	"fmt"
)

type Ball struct {
	posX int
	posY int
	vecX int
	vecY int
}
type Score struct {
	leftPlayer  int
	rightPlayer int
}
type Paddle struct {
	posX int
	posY int
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func drawField(ball *Ball, leftPaddle *Paddle, rightPaddle *Paddle, score *Score) {
	fmt.Printf("\n\t\t\t  LEFT PLAYER %d:%d RIGHT PLAYER\n", score.leftPlayer, score.rightPlayer)
	for i := 0; i <= 25; i++ {
		for j := 0; j <= 80; j++ {
			if i == 0 || i == 25 {
				fmt.Print("-")
			} else if j == 0 || j == 80 {
				fmt.Print("|")
			} else if ball.posX == i && ball.posY == j {
				fmt.Print("@")
			} else if leftPaddle.posX == i && leftPaddle.posY == j {
				fmt.Print("|")
			} else if leftPaddle.posX-1 == i && leftPaddle.posY == j {
				fmt.Print("|")
			} else if leftPaddle.posX+1 == i && leftPaddle.posY == j {
				fmt.Print("|")

			} else if rightPaddle.posX == i && rightPaddle.posY == j {
				fmt.Print("|")
			} else if rightPaddle.posX-1 == i && rightPaddle.posY == j {
				fmt.Print("|")
			} else if rightPaddle.posX+1 == i && rightPaddle.posY == j {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func updateBall(ball *Ball, leftPaddle *Paddle, rightPaddle *Paddle, score *Score) {
	if ball.posX > 23 {
		ball.vecX = -1
	}
	if ball.posX < 2 {
		ball.vecX = 1
	}
	if ball.posY == 76 {
		if ball.posX >= rightPaddle.posX-1 && ball.posX <= rightPaddle.posX+1 {
			ball.vecY = -1
		}
	}
	if ball.posY == 4 {
		if ball.posX >= leftPaddle.posX-1 && ball.posX <= leftPaddle.posX+1 {
			ball.vecY = 1
		}
	}
	if ball.posY > 78 {
		score.leftPlayer++
		ball.posX = 12
		ball.posY = 40
		ball.vecY = -1
	}
	if ball.posY < 2 {
		score.rightPlayer++
		ball.posX = 12
		ball.posY = 40
		ball.vecY = 1
	}
	ball.posX += ball.vecX
	ball.posY += ball.vecY
}

func handleInput(ch rune, leftPaddle *Paddle, rightPaddle *Paddle) {
	switch ch {
	case 'a':
		if leftPaddle.posX > 2 {
			leftPaddle.posX--
		}
	case 'z':
		if leftPaddle.posX < 23 {
			leftPaddle.posX++
		}
	case 'k':
		if rightPaddle.posX > 2 {
			rightPaddle.posX--
		}
	case 'm':
		if rightPaddle.posX < 23 {
			rightPaddle.posX++
		}
	default:
		//
	}
}
func main() {
	ball := &Ball{posX: 12, posY: 40, vecX: 1, vecY: 1}
	leftPaddle := &Paddle{posX: 12, posY: 3}
	rightPaddle := &Paddle{posX: 12, posY: 77}
	score := &Score{leftPlayer: 0, rightPlayer: 0}
	clearScreen()
	drawField(ball, leftPaddle, rightPaddle, score)

	for score.leftPlayer < 5 && score.rightPlayer < 5 {
		var ch rune
		fmt.Scanf("%c", &ch)
		clearScreen()
		handleInput(ch, leftPaddle, rightPaddle)

		updateBall(ball, leftPaddle, rightPaddle, score)
		drawField(ball, leftPaddle, rightPaddle, score)

	}

	fmt.Printf("GAME IS OVER! FINAL SCORE %d:%d", score.leftPlayer, score.rightPlayer)
}
