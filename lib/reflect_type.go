package main

import (
   "fmt"
   "reflect"
)

//反射获取interface类型信息
func reflect_type(x interface{}){
    // 获得Type类型的，Type类型用来表示一个go类型。
   t:= reflect.TypeOf(x)
   // Kind代表Type类型值表示的具体分类。零值表示非法分类。
   k := t.Kind()

   switch k {
   case reflect.Float64 :
      fmt.Println("this is a float64")
   case reflect.Int :
       fmt.Println("this is a int")
   }
}


func main() {
   var x float64 = 3.4
   reflect_type(x)

  var x2 int = 3
  reflect_type(x2)
}