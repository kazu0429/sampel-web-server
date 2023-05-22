package main

import (
     "fmt"
     "log"
     "net/http" // HTTPを扱うパッケージ:クライアントとサーバーを実装する
) 

func helloHandler(w http.ResponseWriter, r *http.Request){
     hello := []byte("Hello, World!!")
     _, err := w.Write(hello)
     if err != nil {
          log.Fatal(err)
     }
}

func main() {
     http.HandleFunc("/hello", helloHandler)
     fmt.Println("Server Start Up.......")
     log.Fatal(http.ListenAndServe("localhost:8080",nil))
}