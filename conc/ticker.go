package main


import (
    "fmt"
    "time"
)

func main(){
    // 1. timer只能响应1次
    /*timer1 := time.NewTimer(time.Second)
    for {
      // 下一次循环会在这里一直阻塞，all goroutines are asleep - deadlock
      <- timer1.C
      fmt.Println("时间到")
    }*/

    // 2. 延时功能
    /*timer2 := time.NewTimer(2*time.Second)
    <-timer2.C
    fmt.Println("2秒到")
    start:=time.Now()
    timerChan :=time.After(2*time.Second)
    <-timerChan
    fmt.Println(time.Since(start))
    fmt.Println("2秒到")
    */

    // 3.停止定时器
    /*timer3 := time.NewTimer(4*time.Second)
    go func (){
        <- timer3.C
        fmt.Println("定时时间到了")
    }()

    time.Sleep(time.Second)
    fmt.Println("1秒钟过去了")
    b :=timer3.Stop()
    if b {
        fmt.Println("计时器已经停止了")
    }*/

   // 4.重置定时器
   /*timer4:=time.NewTimer(3*time.Second)
   timer4.Reset(1*time.Second)
   fmt.Println(time.Now())
   fmt.Println(<-timer4.C)
   fmt.Println("1秒钟过去了")/*

}
 