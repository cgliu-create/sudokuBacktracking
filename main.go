package main

import(
  "fmt"
)
var board = [9][9]int{
    {7,8,0,4,0,0,1,2,0},
    {6,0,0,0,7,5,0,0,9},
    {0,0,0,6,0,1,0,7,8},
    {0,0,7,0,4,0,2,6,0},
    {0,0,1,0,5,0,9,3,0},
    {9,0,4,0,6,0,0,0,5},
    {0,7,0,3,0,0,0,1,2},
    {1,2,0,0,0,7,4,0,0},
    {0,4,9,2,0,6,0,0,7},
  }
func printBoard(board [9][9]int) {
  fmt.Println("-----------")
  for i, slice := range board {
    for j, num := range slice {
      fmt.Print(num)
      if j==2 || j==5 {
        fmt.Print("|")
      }
      if j==8 {
        fmt.Println()
      }
    }
    if i==2 || i==5 {
      fmt.Println("-----------")
    }
  }
  fmt.Println("-----------")
}
func findEmpty(board [9][9]int) (int, int) {
  for r := 0; r < 9; r++ {
    for c := 0; c < 9; c++ {
       if board[r][c] == 0 {
         return r, c
       }
     }
   }
   return -1, -1
}
func checkValid(row, col, num int, board[9][9]int) bool {
  for i := 0; i < 9; i++ {
    if board[row][i]==num && col != i {
      return false
    }
  }
  for i := 0; i < 9; i++ {
    if board[i][col]==num && row != i{
      return false
    }
  }
  startRow := row / 3
  startCol := col / 3
  for r := startRow*3; r < startRow*3+3; r++ {
    for c := startCol*3; c < startCol*3+3; c++ {
      if board[r][c]==num && r!=row && c!=col {
        return false
      }
    }
  }
  return true;
}

func solveSudoku(board [9][9]int) bool {
  r, c := findEmpty(board)
  if r==-1 && c==-1 {
    return true
  }
  for i:= 1; i <= 9; i++ {
    if checkValid(r, c, i, board) {
      board[r][c] = i
      if solveSudoku(board) {
        printBoard(board)
        return true
      }
      board[r][c] = 0
    }
  }
  return false
}
func main() {
  printBoard(board)
  solveSudoku(board)
}

