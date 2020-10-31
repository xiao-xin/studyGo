package main

import  (
    "fmt"
    "net"
    "bufio"
    "io"
)

func main (){
    // 1. 监听tcp协议，8080顿口
     listen, err :=net.Listen("tcp",":8080")
     if err != nil {
        fmt.Printf("failed listen,err:%s",err)
        return
     }
     defer listen.Close()
    // 2. 接收客户端的连接请求
    for {
        conn,err := listen.Accept()
        if err != nil {
          fmt.Printf("faield listen,err:%s",err)
             continue
        }
        // 3. 开启协程处理请求
        go process(conn)
    }
}

func process(conn net.Conn){
    defer conn.Close()
    // 需要循环处理来自客户端发送的数据
    for {
        // 使用buffer byte 接收客户端的请求
        buf := make([]byte,512)
        reader := bufio.NewReader(conn)
        n, err := reader.Read(buf)
        if err != nil && err != io.EOF {
            fmt.Printf("received message faile ,err:%s\n",err)
            break
        }

        recvStr := string(buf[:n])
        fmt.Printf("received message from client,message:%s\n",recvStr)
        // 响应到客户端
        conn.Write(buf[:n])
    }
}














/*
func main(){
    listen,err := net.Listen("tcp","127.0.0.1:8080")
    if err != nil {
        fmt.Printf("failed listen, err:%s/n",err)
        return
    }
    defer listen.Close()
    fmt.Printf("listen tcp\n")
    for {
         conn, err := listen.Accept() // 建立连接
         if err != nil {
          fmt.Printf("failed connection, err:%s/n",err)
          continue
         }
         go process(conn)
    }
}

func process(conn net.Conn){
    defer conn.Close()

    for {
        reader := bufio.NewReader(conn)
        var buf [128]byte

        n, err := reader.Read(buf[:])// 读取数据

        if err != nil {
            break;
        }
        recvStr := string(buf[:n])
        fmt.Printf("received message from client,message:%s\n",recvStr)
        conn.Write(buf[:n])
    }
}*/

