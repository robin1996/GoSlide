package main

import (
	"fmt"
	"time"
	"math/rand"
	
	"github.com/nsf/termbox-go"
)

// game move
const (
	up = iota // 0
	left // 1
	right // 2
	down // 3
)

var blankXPos int = 0
var blankYPos int = 0

var completedBoard = board{{'0', '1', '2'},
						   {'3', '4', '5'},
						   {'6', '7', '8'}}

var gameBoard = completedBoard

type tile = rune

type board = [3][3]tile

// Also sets blank pos vars
func drawBoard(xPos, yPos int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 9; x++ {
			termbox.SetCell(xPos + x, yPos + y, ' ', termbox.ColorDefault, 0xC618)
		}
	}
	for y, l := range gameBoard {
		for x, t := range l {
			if t == completedBoard[0][0] {
				termbox.SetCell(xPos + 2 + 2 * x, yPos + 1 + y, ' ', termbox.ColorBlack, 0xC618)
				blankYPos = y
				blankXPos = x
			} else {
				termbox.SetCell(xPos + 2 + 2 * x, yPos + 1 + y, t, termbox.ColorBlack, 0xC618)
			}
		}
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawBoard(2, 1)
	termbox.Flush()
}

func checkMoveLegal(move *int) bool {
	switch *move {
	case up:
		return blankYPos < 2
	case left:
		return blankXPos < 2
	case right:
		return blankXPos > 0
	case down:
		return blankYPos > 0
	}
	return false
}

func slide(move *int) {
	switch *move {
	case up:
		gameBoard[blankYPos][blankXPos], gameBoard[blankYPos + 1][blankXPos] = gameBoard[blankYPos + 1][blankXPos], gameBoard[blankYPos][blankXPos]
	case down:
		gameBoard[blankYPos][blankXPos], gameBoard[blankYPos - 1][blankXPos] = gameBoard[blankYPos - 1][blankXPos], gameBoard[blankYPos][blankXPos]
	case right:
		gameBoard[blankYPos][blankXPos], gameBoard[blankYPos][blankXPos - 1] = gameBoard[blankYPos][blankXPos - 1], gameBoard[blankYPos][blankXPos]
	case left:
		gameBoard[blankYPos][blankXPos], gameBoard[blankYPos][blankXPos + 1] = gameBoard[blankYPos][blankXPos + 1], gameBoard[blankYPos][blankXPos]
	}
}

func playerMove(move int) {
	if checkMoveLegal(&move) {
		slide(&move)
	} else {
		// Illegal Move
	}
}

func shuffle() {
	for i := 0; i < 20; {
		move := rand.Intn(3)
		if checkMoveLegal(&move) {
			slide(&move)
			draw()
			time.Sleep(300 * time.Millisecond)
			i++
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	shuffle()
	for {
		draw()
		keyPress := termbox.PollEvent()
		switch keyPress.Type {
		case termbox.EventKey:
			switch keyPress.Key {
			case termbox.KeyArrowUp:
				playerMove(up)
			case termbox.KeyArrowDown:
				playerMove(down)
			case termbox.KeyArrowLeft:
				playerMove(left)
			case termbox.KeyArrowRight:
				playerMove(right)
			case termbox.KeyEsc:
				return
			case termbox.KeyCtrlC:
				return
			default:
				fmt.Println("other")
			}
		case termbox.EventError:
			panic(keyPress.Err)
			}
		}
}
