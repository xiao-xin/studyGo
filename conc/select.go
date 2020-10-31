package main


import (
    "fmt"
    "time"
)


/*
    select

   我们也可以从for中获取多个通道的值，销量比较低
   for{
       // 尝试从ch1接收值
       data, ok := <-ch1
       // 尝试从ch2接收值
       data, ok := <-ch2
       …
   }

   我们可以使用select操作多个通道，发送、接收通道中的数据都是可以的。

   他的工作方式是这样的：
   select {
       case <-chan1:
          // 如果chan1成功读到数据，则进行该case处理语句
       case chan2 <- 1:
          // 如果成功向chan2写入数据，则进行该case处理语句
       default:
          // 如果上面都没有成功，则进入default处理流程
   }
*/
// select可以同时监听一个或多个channel，直到其中一个channel ready
func demo1 (){
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func(ch chan string ){
    time.Sleep(2*time.Second)
        ch <- "i am chan 1"
    }(ch1)

    go func(ch chan string ){
        time.Sleep(5*time.Second)
        ch <- "i am chan 2"
    }(ch2)

    select {
        case ret :=  <- ch1 :
             fmt.Println(ret)
        case ret :=  <- ch2 :
             fmt.Println(ret)
    }
}

// 如果多个channel同时ready，则随机选择一个执行
func demo2(){
     // 创建2个管道
       int_chan := make(chan int, 1)
       string_chan := make(chan string, 1)
       go func() {
          //time.Sleep(2 * time.Second)
          int_chan <- 1
       }()
       go func() {
          string_chan <- "hello"
       }()

       time.Sleep(time.Second)
       select {
       case value := <-int_chan:
          fmt.Println("int:", value)
       case value := <-string_chan:
          fmt.Println("string:", value)
       }
       fmt.Println("main结束")
}

// 可以用select判断管道是否存满
func demo3 (){
    job := make(chan int,10)

    go func(ch chan<- int){
        for i:=1;;i++ {
           select {
             case  ch <-i :
                fmt.Println("write",i)
             default :
                fmt.Println("chan full")
           }
           time.Sleep(100*time.Millisecond )
        }
    }(job)

    for v := range job {
        fmt.Println("recv:",v)
        time.Sleep(2*time.Second)
    }
}

func main() {
    demo3()
}

