/**
 *
 * 功能：简易路由器
 * 作者: 吴加亮
 * 创建时间:2019年7月19日13:03:49
*/

package main

import(
  "fmt"
  "net/http"
)

type MyMux struct{
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request){
    if r.URL.Path=="/"{
        sayHelloName(w,r)
        return 
    }
    http.NotFound(w,r)
    return
}

func sayHelloName(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"Hello World!")//这个写入到W的是输出到客户端的
}

func main(){
   mux:=&MyMux{}
   http.ListenAndServe(":9090",mux)
   fmt.Println("http://localhost:9090")
}