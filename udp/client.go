package main

import  (
    "fmt"
    "net"
    //"bufio"
    //"io"
)

func main (){
    conn , err := net.DialUDP("udp",nil,&net.UDPAddr{
        IP : net.IPv4(127,0,0,1),
        Port:8080,
    })
     if err != nil {
            fmt.Printf("conn failed,err:%v\n",err)
            return
     }
    defer conn.Close()

    sendData := []byte("hell upd")
    _,err =conn.Write(sendData)
    if err != nil {
        fmt.Printf("send failed,err:%v\n",err)
        return
    }

    recvData:= []byte("hell upd")
    n,err :=conn.Read(recvData)
      if err != nil {
            fmt.Printf("recv failed,err:%v\n",err)
            return
      }
      fmt.Printf("rece,msg:%s",recvData[:n])











    /*client, err := net.DialUDP("udp",nil,&net.UDPAddr{
        IP:net.IPv4(127,0,0,1),
        Port:8080,
    })
     if err != nil {
       fmt.Printf("conn failed,err:%v\n",err)
       return
     }
     defer client.Close()

     sendData := []byte("hello go")
     _,err =client.Write(sendData)
     if err != nil {
       fmt.Printf("send failed,err:%v\n",err)
       return
     }
     receData := make([]byte,1024)
     n,addr,err :=client.ReadFromUDP(receData)
     fmt.Printf("recv data:%s,add:%v,size:%v\n",receData[:n],addr,n)
     */
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

