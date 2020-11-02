package main

import (
   "fmt"
   "reflect"
)

//反射设置interface类型的值
func reflect_set_value(x interface{}){
   
   v:= reflect.ValueOf(x)
   // Kind返回v持有的值的分类，如果v是Value零值，返回值为Invalid
   k := v.Kind()
   if k != reflect.Ptr {
      return 
   }

   // Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
   t := v.Elem();

   switch t.Kind() {
   case reflect.Float64 :
      fmt.Printf("set before value:%f\n",t.Float())
      t.SetFloat(3.1415)
      fmt.Println("set after value:%f",t.Float())
      // 地址
      // 将v持有的值作为一个指针返回。本方法返回值不是unsafe.Pointer类型
      fmt.Println(v.Pointer())
}
}


func main() {
   var x float64 = 3.4
   reflect_set_value(&x)
    fmt.Printf("%p",&x)
}