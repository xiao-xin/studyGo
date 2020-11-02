package main

import (
   "fmt"
   "reflect"
)

//反射获取interface类型信息
func reflect_type(x interface{}){
   k ,err := reflect.TypeOf(x)
   type := k.Kind()

   switch type{
   case reflect.Float64 :
      fmt.Println("this is a float64")
   }
}


func main() {
   var x float64 = 3.4
   reflect_type(x)
}