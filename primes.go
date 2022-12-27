// A concurrent prime sieve.

package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
  for i := 2; ; i++ {
    ch <- i
  }
}

// Copy values from 'in' to 'out', removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
  for {
    i := <-in
    if i%prime != 0 {
      out <- i
    }
  }
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
  ch := make(chan int)
  go Generate(ch)
  for i := 0; i < 10; i++ {
    prime := <-ch
    fmt.Println(prime)
    ch1 := make(chan int)
    go Filter(ch, ch1, prime)
    ch = ch1
  }
}
