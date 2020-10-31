package main


import (
    "fmt"
    "time"
    "sync/atomic"
    "sync"
)

/*
   传统的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。
   对于基本数据类型可以使用原子操作来保证并发安全，性能比加锁要好

  sync.atomic包提供了底层原子级内存的操作方法，对于同步算法的实现很有用。
*/

var x int64
var lock sync.Mutex
var wg sync.WaitGroup

// 普通版本的函数，并发不安全
func add (){
  x ++
  wg.Done()
}

// 加锁版本，并发安全，开销大
func mutexAdd(){
  lock.Lock()
  x ++
  lock.Unlock()
  wg.Done()
}

// 原子操作并发安全
func atomicAdd(){
  atomic.AddInt64(&x,1)
  wg.Done()
}

func main(){
    
  start := time.Now()
  for i:=0;i<10000000;i++{
    wg.Add(1)
    //go  add()
    //go mutexAdd()
    go atomicAdd()
  }

  wg.Wait()
  fmt.Println(time.Since(start))
  fmt.Println(x)
}
 