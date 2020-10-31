package main

import  (
    "fmt"
    "net"
    "bufio"
    "os"
    "strings"
)
/*func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Println("err :", err)
        return
    }
    defer conn.Close() // 关闭连接
    inputReader := bufio.NewReader(os.Stdin)
    for {
        input, _ := inputReader.ReadString('\n') // 读取用户输入
        inputInfo := strings.Trim(input, "\r\n")
        if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
            return
        }
        _, err = conn.Write([]byte(inputInfo)) // 发送数据
        if err != nil {
            return
        }
        buf := [512]byte{}
        n, err := conn.Read(buf[:])
        if err != nil {
            fmt.Println("recv failed, err:", err)
            return
        }
        fmt.Println(string(buf[:n]))
    }
}*/


func main (){
    // 1. 连接tcp服务器
    conn, err := net.Dial("tcp","127.0.0.1:8080")
    if err != nil {
        fmt.Printf("conn tcp server failed,err:%s",err)
        return
    }
    defer conn.Close()

     reader := bufio.NewReader(os.Stdin)
    for {

         // 2. 准备发送的内容
         input,err:= reader.ReadString('\n')
         if err != nil {
             fmt.Printf("read string ,err:%s",err)
         }
         inputInfo := strings.Trim(input,"\r\n")
           if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
                      return
                  }
          // 3. 发送内容到服务器
         _,err=conn.Write([]byte(inputInfo))
         if err != nil {
            fmt.Printf("send message failed,err:%s",err)
            return
         }
          // 4. 接收服务器端的响应
         msg := make([]byte,512)
         n,err:=conn.Read(msg)
         if err != nil {
            fmt.Printf("received message failed,err:%s",err)
            return
         }
         fmt.Printf("received server message:%s",msg[:n])
    }
}















/*func main(){
    conn, err := net.Dial("tcp","127.0.0.1:8080")
    if err != nil {
        fmt.Printf("connection tcp server failed,err:%s\n",err)
        return
    }
    defer conn.Close()

    buf :=  bufio.NewReader(os.Stdin)
    for {
        input , _ := buf.ReadString('\n')
        inputInfo :=strings.Trim(input,"\r\n")
        if strings.ToUpper(inputInfo) == "Q" {
            return
        }

        _, err = conn.Write([]byte(inputInfo))
        if err != nil {
         fmt.Printf("send message failed,err:%s\n",err)
          return
        }

        var msg [512]byte
        n, err := conn.Read(msg[:])
        if err != nil {
            fmt.Printf("recv failed,err:%s\n",err)
            return
        }
        fmt.Printf("recv message:%s\n",msg[:n])
    }
}

*/

