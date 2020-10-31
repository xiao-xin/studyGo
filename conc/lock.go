package main


import (
    "fmt"
    "sync"
    "time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex
/*
1.并发安全

一般对于几个goroutine同时访问资源，会发生竞态问题，一般常用的做法就是加锁。
只有获得资源的锁时才可以对资源进行操作，其他的goroutine只能等待。
*/

// 案例1
// 下面这个列子就是不加锁的情况，每次得到x的结果都是不一样的。

func add1 (){
	for i :=0 ; i< 5000 ; i ++ {
		x = x+ int64(i)	
	}
	wg.Done()
}
func demo1 (){
  wg.Add(2)
  go add1()
  go add1()
  wg.Wait()
  fmt.Println(x)
}

/*
	2.互斥锁
	常用的做法就是使用互斥锁来控制共享资源访问，能够保证只有一个goroutine访问共享资源，
	使用sync包的Mutex类型来实现互斥锁

	互斥锁可以保证同一个时间只有一个goroutine进入临界区，其他的goroutine只能等待锁，
	当互斥锁被释放后其他正在等待的goroutine才可以临界区。 如果有多个goroutine等待一个锁，
	唤醒哪个goroutine是随机的。
*/

// 案例二
// 这种结果后每次运行后都一样
func add2 (){
	for i :=0 ; i< 5000 ; i ++ {
		lock.Lock()  // 加锁
		x = x+ int64(i)	
		lock.Unlock()  // 解锁
	}
	wg.Done()
}
func demo2 (){
  wg.Add(2)
  go add2()
  go add2()
  wg.Wait()
  fmt.Println(x)
}

/*
	3.读写互斥锁

	上面的互斥锁是完全互斥的，也就说读也是互斥的，在读多写少的场景下，
	并发的去读取一个资源不涉及资源的修改是没有必要加锁的，这种情况使用
	读写锁会更好，可以使用sync包中的RWMutex类型添加读写锁

	读写锁有读锁和写锁，当一个goroutine获得读锁后，其他goroutine如果是获取
	读锁会继续获得不会等待，如果是获取写锁就需要等待了。

	如果一个goroutine获得写锁，其他的goroutine不管是获取读锁还是写锁都会等待
*/


// 案例三
// 这里输出的结果都是1，也可能不是，总之在有读锁时，获得写锁的goroutine只能等待。
func read (){
	rwlock.RLock() // 获得读锁
	time.Sleep(5*time.Millisecond)  // 假设读操作耗时1毫秒
	fmt.Println(x)
	rwlock.RUnlock()  // 解读锁
	wg.Done()
}

func write (){
	rwlock.Lock()  // 获取写锁
	x=x+1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()  // 解写锁
	wg.Done()
}

func demo3 (){
 	for i:=0 ; i<10 ; i++{
 		wg.Add(1)
 		go write()
 	}

 	for i:=0 ; i<100 ; i++{
 		wg.Add(1)
 		go read()
 	}

 	wg.Wait()
}

func main(){
	demo3()
}