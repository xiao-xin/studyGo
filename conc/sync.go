package main


import (
    "fmt" 
    "sync"
    "strconv"
)


/*

 1.并发同步

 并发同步用Sleep来做肯定是不合适的，使用sync包中的WaitGroup来实现并发任务的
 同步。

 sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
 例如当我们启动了N 个并发任务时，就将计数器值增加N。
 每个任务完成时通过调用Done()方法将计数器减1。
 通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

 注意sync.WaitGroup是一个结构体，传递的时候要传递指针。
*/

/*var wg sync.WaitGroup
func demo1(){
  wg.Add(1)

  go  func (wg *sync.WaitGroup){
    defer wg.Done()
    fmt.Println("goroutine end")
  }(&wg)

  wg.Wait()
}*/

/*
  2. sync.Once()

  在有些场景下比如初始化操作在并发执行时值需要初始化一次，只关闭一次通道。
  使用sync.Once中的Do方法。

  func (o *Once) Do(f func()) {}


  延迟加载初始化在操作在用到它的时候再执行是一个很好的策略，这样可以减少程序的启动耗时，
  有可能在实际的执行中都没有用上这个初始化变量。

  多个协程初始化操作会被执行多次，几遍我们添加了判断也会这样，除非我们添加互斥锁，这样
  一来又会引发性能。 同时还可能引发数据安全的问题。
   func LoadConfig(){
     config = map[string]string{
      "name":"jack",
      "age":"25",
     }
     fmt.Println(config)
   }

   func Load(){
     if config == nil {
      LoadConfig()
     }
     wg.Done()
   }

  看到上面多个协程调用Load函数时不是并发安全的，这是因为CPU为了保证每个goroutine
  都满足串行一致会自动的重排内存的顺序。LoadConfig函数可能会被重排为以下结果：
  func LoadConfig(){
     config = make(map[string]string)
     config["name"] = "jack"
     config["age"] = "25"
   }

  这样一来通过config不是nil并不意味着初始化已经完成了，可以加互斥锁，但是效率低下。
  */

 /*var config map[string]string

 func LoadConfig(){
   config = map[string]string{
    "name":"jack",
    "age":"25",
   }
   fmt.Println(config)
 }

 func Load(){
   if config == nil {
    LoadConfig()
   }
   wg.Done()
 }

 func main(){
    
    for i:=0;i<100;i++{
      wg.Add(1)
      go Load()
    }
    wg.Wait()
 }*/



  /*
  可以使用sync.Once来运行LoadConfig函数，其实在Once内部包含一个互斥锁和布尔值，互斥锁
  保证数据和布尔值是安全的，布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候
  是并发安全的并且初始化操作也不会被执行多次。
*/

/* var config map[string]string
 var loadConfigOnce  sync.Once

 func LoadConfig(){
   config = map[string]string{
    "name":"jack",
    "age":"25",
   }
   fmt.Println(config)
 }

 func Load(){
   loadConfigOnce.Do(LoadConfig)
   wg.Done()
 }

 func main(){
    
    for i:=0;i<100;i++{
      wg.Add(1)
      go Load()
    }
    wg.Wait()
 }*/



/*
  3.内置的Map也不是并发安全的,比如如下这个例子。会引发fatal error: concurrent map writes
*/
/*
var m map[string]int

func get(key string ) int {
  return m[key] 
}

func set(key string, v int){
  m[key]  = v
}

func main (){
  m  = make(map[string]int)

  for i:=0;i<100;i++{
    wg.Add(1)
    go func (v int){
      key := strconv.Itoa(v)
      set(key,i)
      fmt.Println("k:%v,v:%d",key,get(key))
      wg.Done()
    }(i)
  }
  wg.Wait()
}*/

/*
  3.需要对map加锁来保证并发的安全性，提供了一个sync.Map，开箱急用不需要使用make初始化
  并在内部提供了Store、Load、LoadOrStore、Delete、Range等操作方法
*/

var m sync.Map
func main (){

  for i:=0;i<100;i++{
    wg.Add(1)
    go func (v int){
      key := strconv.Itoa(v)
      m.Store(key,v)
      keyV,_:=m.Load(key)
      fmt.Printf("k:%v,v:%d\n",key,keyV)
      wg.Done()
    }(i)
  }
  wg.Wait()
}