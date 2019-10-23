package main

import (
  "math/rand"
  "sync/atomic"
  "time"
  "fmt"
  "sync"
)

const MAX_THREAD, MAX_LENGTH = 5, 20

func main() {
  rand.Seed(time.Now().Unix())
  n := (rand.Int() + 100) % 1001
  // generate array
  array := make([]int, n)
  for i := 0; i < n; i++ {
    array[i] = rand.Int() % 20
  }
  fmt.Println(add(array))

  result := 0
  for i := 0; i < n; i++ {
    result += array[i]
  }
  fmt.Println(result)
}

func add(arr []int) int {
  pool := make(chan int, MAX_THREAD)
  for i := 0; i < MAX_THREAD; i++ {
    pool <- 1
  }
  
  i, n := 0, len(arr)-1
  var result int32
  var wg sync.WaitGroup
  for {
    <-pool
    wg.Add(1)

    if i >= n {
      break
    }

    var sub_arr []int
    if i + MAX_LENGTH >= n {
      sub_arr = arr[i:]
      i = n
    } else {
      sub_arr = arr[i:i+MAX_LENGTH]
      i = i + MAX_LENGTH
    }

    go func(sub_arr []int) {
      sum := 0
      for _, num := range sub_arr {
        sum += num
      }
      atomic.AddInt32(&result, int32(sum))

      pool <- 1
      wg.Done()
    }(sub_arr)
  }

  wg.Wait()
  return int(result)
}