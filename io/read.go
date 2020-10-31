package main

import  (
    "fmt"

       "io"
       "strings"
)

func readForm(read io.Reader,n int) ([]byte,error){
    buf := make([]byte,n)
    n,err :=read.Read(buf)
    if n >0 {
        return buf[:n],nil
    }
    return nil,err
}

func main(){
//     data , _:=readForm(os.Stdin,2)
//     fmt.Printf("data:%s",data)

    reader := strings.NewReader("hello")
    data , _:=readForm(reader,2)
    fmt.Printf("data:%s",data)
}



