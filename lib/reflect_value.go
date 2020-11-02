package main

import (
   "fmt"
   "reflect"
)

//反射获取interface类型的值
func reflect_value(x interface{}){
   // 返回一个初始化为i接口保管的具体值的Value，ValueOf(nil)返回Value零值。
   v:= reflect.ValueOf(x)

   // Kind返回v持有的值的分类，如果v是Value零值，返回值为Invalid
   k := v.Kind()

   switch k {
   case reflect.Float64 :
      fmt.Println(v.Float())
   case reflect.Int :
       fmt.Println(v.Int())
   }
}


func main() {
   var x float64 = 3.4
   reflect_value(x)

  var x2 int = 3
  reflect_value(x2)
}