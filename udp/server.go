package main

import  (
    "fmt"
    "net"
    //"bufio"
    //"io"
)

func main (){
    listen , err := net.ListenUDP("udp",&net.UDPAddr{
       IP : net.IPv4(0,0,0,0),
       Port:8080,
    })
   if err != nil {
        fmt.Printf("listen failed,err:%v\n",err)
        return
    }
    defer listen.Close()

    for {
        buf := make([]byte,1024)
        n, addr, err := listen.ReadFromUDP(buf)
        if err != nil {
                fmt.Printf("rece msg failed,err:%v\n",err)
                continue
        }
        fmt.Printf("data:%s,addr:%s,size:%d\n",buf[:n],addr,n)
        listen.WriteToUDP(buf[:n],addr)
    }









     /*
    listen, err := net.ListenUDP("udp",&net.UDPAddr{
        IP: net.IPv4(0,0,0,0),
        Port:8080,
    })
    if err != nil {
        fmt.Printf("listen failed,err:%v\n",err)
        return
    }
    defer listen.Close()

    for {
         msg := make([]byte,1024)
         n,addr,err := listen.ReadFromUDP(msg)
         if err != nil {
             fmt.Printf("listen failed,err:%v\n",err)
             continue
         }

         fmt.Printf("data:%s,add:%v,size:%v\n",msg[:n],addr,n)
         sendMsg := "server:"+string(msg[:n])
         n ,err = listen.WriteToUDP([]byte(sendMsg),addr)
         if err != nil {
            fmt.Printf("write message faield ,err:%v\n",err)
        }
    }*/
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

