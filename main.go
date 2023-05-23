package main

import (
     "fmt"
     "os"
     "bufio"
     "log"
     "strings"
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
     for _, elm := range contentlist{
          fmt.Println(strings.Split(elm, " "))
     }
     html, err := template.ParseFiles("./html/go_playground.html")
     if err != nil {
          log.Println(err)
     }
     getContens := New(contentlist)
     // htmlファイルにcontens情報を渡している
     if err := html.Execute(w, getContens); err != nil{
          log.Println(err)
     }
}

/* 
os.O_WRONLY 書き込み専用, 
os.O_APPEND 追記,
os.O_CREATE 作成
引用 https://golang.hateblo.jp/entry/2018/11/09/163000

os.FileMode(0600) パーミッション
*/ 
func createHandler(w http.ResponseWriter, r *http.Request){
     formValue := r.FormValue("value")
     file, err := os.OpenFile("go_playground.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
     defer file.Close()
     if err != nil{
          log.Println(err)
     }
     _, err = fmt.Fprintln(file, formValue)
     if err != nil {
          log.Println(err)
     }
     http.Redirect(w, r, "/go", http.StatusFound)
}

/*
http.ListenAndServe サーバー起動
 - arg1 tcpアドレス
 - http.Handler 

引用 https://journal.lampetty.net/entry/understanding-http-handler-in-go
*/

func main() {
     http.HandleFunc("/go", viewHandler)
     http.HandleFunc("/go/create", createHandler) // actionと同じpath
     fmt.Println("localhost:8080")
     fmt.Println("Server Start Up.......")
     log.Println(http.ListenAndServe("localhost:8080",nil))
}