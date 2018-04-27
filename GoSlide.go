package main

import (
	"fmt"
	
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
var tiles = [9]tile{0, 1, 2, 3, 4, 5, 6, 7, 8}
var gameBoard = board{{&tiles[0], &tiles[1], &tiles[2]}, 
					  {&tiles[3], &tiles[4], &tiles[5]},
					  {&tiles[6], &tiles[7], &tiles[8]}}

type tile = int

type square = *tile

type board = [3][3]square

// Also sets blank pos vars
func drawBoard() {
	fmt.Println("+---------+")
	for y, l := range gameBoard {
		fmt.Print("|")
		for x, s := range l {
			fmt.Print(" ")
			if *s == 0 {
				fmt.Print(" ")
				blankYPos = y
				blankXPos = x
			} else {
				fmt.Print(*s)
			}
			fmt.Print(" ")
		}
		fmt.Print("|\n")
	}
	fmt.Println("+---------+")
}

func checkMoveLegal(move int) bool {
	switch move {
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

func slide(move int) {
	switch move {
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
	if checkMoveLegal(move) {
		slide(move)
	} else {
		fmt.Println("Illegal move!")
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	for {
		drawBoard()
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
