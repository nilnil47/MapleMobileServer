package main

import (
	"fmt"
	"testing"
)

type tool struct {
	movement int
}

func (b tool) getMovement() int {
	return b.movement
}

type boardTool interface {
	getMovement() int
}

func TestA(t *testing.T) {
	bishop := tool{3}
	knight := tool{5}
	board := [10][10]*tool{}
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if x > y {
				board[x][y] = &bishop
			} else {
				board[x][y] = &knight
			}
		}
	}

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Print(board[x][y].getMovement())
		}
		fmt.Println()
	}
}

//func TestMain(t *testing.T) {
//	bishop := tool{3}
//	knight := tool{5}
//	board := [10][10]*tool{}
//	for x := 0; x < 10; x++ {
//		for y := 0; y < 10; y++ {
//			if x > y {
//				board[x][y] = &bishop
//			} else {
//				board[x][y] = &knight
//			}
//		}
//	}
//
//	for x := 0; x < 10; x++ {
//		for y := 0; y < 10; y++ {
//			fmt.Print(board[x][y].getMovement())
//		}
//		fmt.Println()
//	}
//
//}
