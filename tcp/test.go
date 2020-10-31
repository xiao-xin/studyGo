package main

import  (
    "fmt"
    "os"
    "bufio"
    "io"
)

func main(){
    rd := bufio.NewReader(os.Stdin)  // f is io.Reader interface.
    for{
       line, err := rd.ReadString('\n')
        if err != nil || io.EOF ==err{
           break
       }
       fmt.Printf("%s#",line)
    }
}
