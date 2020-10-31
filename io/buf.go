package main

import  (
    "fmt"

       "os"
       //"bufio"
)

func main(){

     // 1.创建文件
     filename := "./test.txt"
     // 写入
     file,err:=os.OpenFile(filename,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
     if err != nil {
        fmt.Printf("open file failed,err:%v",err)
        return
     }

     file.WriteString("hello go")
     file.WriteAt([]byte("java"),6)

      file, err := os.Open(filename)
      if err != nil {
             fmt.Printf("open file failed,err:%v",err)
             return
      }
      by := make([]byte,2)
       // 读取
      n,err:=file.ReadAt(by,2)
      fmt.Printf("%s",by[:n])
}



