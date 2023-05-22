package main

import (
     "fmt"
     "log"
     "html/template"
     "net/http" // HTTPを扱うパッケージ:クライアントとサーバーを実装する
)

func viewHandler(w http.ResponseWriter, r *http.Request){
     html, err := template.ParseFiles("go_playground.html")
     if err != nil {
          log.Fatal(err)
     }
     if err := html.Execute(w, nil); err != nil{
          log.Fatal(err)
     }
}

func main() {
     http.HandleFunc("/go", viewHandler)
     fmt.Println("localhost:8080")
     fmt.Println("Server Start Up.......")
     log.Fatal(http.ListenAndServe("localhost:8080",nil))
}