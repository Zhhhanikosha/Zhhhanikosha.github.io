// Shugayeva Zhaniya
// Write a program that takes a list of integers as input and sorts the list in ascending order.
package main

import (
 "fmt"
)

func bubbleSort(arr []int) {
 n := len(arr)

 for i := 0; i < n-1; i++ {
  for j := 0; j < n-i-1; j++ {
   if arr[j] > arr[j+1] {
    arr[j], arr[j+1] = arr[j+1], arr[j]
   }
  }
 }
}

func main() {
 arr := []int{9, 2, 7, 4, 5, 1, 3, 8, 6}

 bubbleSort(arr)

 fmt.Println(arr)
}
