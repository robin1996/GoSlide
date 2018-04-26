package main

import (
	"fmt"
	
	"github.com/nsf/termbox-go"
)

type tile = int

type board = [3][3]*tile

func drawBoardLine(boardLine [3]*tile) {
	fmt.Print("|")
	for _, t := range boardLine {
		fmt.Print(" ")
		if *t == 0 {
			fmt.Print(" ")
		} else {
			fmt.Print(*t)
		}
		fmt.Print(" ")
	}
	fmt.Print("|\n")
}

func drawBoard(gameBoard board) {
	fmt.Println("+---------+")
	for _, l := range gameBoard {
		drawBoardLine(l)
	}
	fmt.Println("+---------+")
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	var tiles = [9]tile{0, 1, 2, 3, 4, 5, 6, 7, 8}
	var testBoard = board{{&tiles[0], &tiles[1], &tiles[2]}, {&tiles[3], &tiles[4], &tiles[5]}, {&tiles[6], &tiles[7], &tiles[8]}}
	for {
		keyPress := termbox.PollEvent()
		switch keyPress.Type {
		case termbox.EventKey:
			switch keyPress.Key {
			case termbox.KeyArrowUp:
				fmt.Println("up")
			case termbox.KeyArrowDown:
				fmt.Println("down")
			case termbox.KeyArrowLeft:
				fmt.Println("left")
			case termbox.KeyArrowRight:
				fmt.Println("right")
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
		drawBoard(testBoard)
}
