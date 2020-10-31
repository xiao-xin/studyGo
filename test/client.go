package main

import (
    "fmt"

)



func one (ch chan<- int){
   for i := 0; i<100 ; i ++{
         ch <- i
    }
    close(ch)
}

func two (ch1  <-chan int,ch2  chan<- int){
    for i := range ch1{
               ch2 <- i*i
    }
    close(ch2)
}

func three (ch  <-chan int){
    for i := range ch{
       fmt.Println(i)
    }
}

func main() {
   ch1 := make(chan int)
   ch2 := make(chan int)

   go one(ch1)
   go two(ch1,ch2)
   three(ch2)
}