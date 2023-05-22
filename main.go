package main

import (
     "fmt"
     "os"
     "bufio"
     "log"
     "html/template"
     "net/http" // HTTPを扱うパッケージ:クライアントとサーバーを実装する
)

type ContentList struct {
     Contents []string
}

func New(contents []string) *ContentList{
     return &ContentList{Contents: contents}
}

func fileRead(fileName string) []string{
     var contentlist []string
     file, err := os.Open(fileName) // For read access.
     if os.IsNotExist(err){
          return nil
     }
     defer file.Close() // 非同期
     scaner := bufio.NewScanner(file)
     for scaner.Scan() {
          contentlist = append(contentlist, scaner.Text())
     }
     return contentlist
}

func viewHandler(w http.ResponseWriter, r *http.Request){
     contentlist := fileRead("go_playground.txt")
     fmt.Println(contentlist)
     html, err := template.ParseFiles("go_playground.html")
     if err != nil {
          log.Fatal(err)
     }
     getContens := New(contentlist)
     // htmlファイルにcontens情報を渡している
     if err := html.Execute(w, getContens); err != nil{
          log.Fatal(err)
     }
}

func main() {
     http.HandleFunc("/go", viewHandler)
     fmt.Println("localhost:8080")
     fmt.Println("Server Start Up.......")
     log.Fatal(http.ListenAndServe("localhost:8080",nil))
}