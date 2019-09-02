package main

import (
  "fmt"
)

func main() {
  
  
  b := [15]int{0, 0, 2, 0, 0, 6, 0, 0, 0, 4}

  smallerNo := -1
  countOfZero := 0
  totalUnitOfWater := 0
  lastNonZeroNumber := -1
  for _ , v := range b {
    if v!= 0 {
    
      if smallerNo == -1 {
	smallerNo = v
      } else {
	smallerNo = min(smallerNo, v)
      }
      lastNonZeroNumber = v

      if (smallerNo > -1 && countOfZero > 0) {
	mul := countOfZero * smallerNo
	totalUnitOfWater = totalUnitOfWater + mul
	smallerNo = lastNonZeroNumber
	countOfZero = 0
      }
    }

    if v == 0 && smallerNo > -1 {
      countOfZero = countOfZero + 1
    }

  }

  fmt.Print("Units of water stored: ", totalUnitOfWater)

}

func min(a, b int) int {
  if a <= b {
    return a
  }
  return b
}

  

