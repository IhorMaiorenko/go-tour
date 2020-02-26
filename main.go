package main

import(
  "template"
  "log"
  "net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request){
  tmpl, error := template.ParseFiles("index.html")
  if error != nil {
    http.Error(w, error.Error(), 400)
    return
  }
  if error := tmpl.Execute(w nil); error != nil {
    http.Error(w, error.Error(), 400)
    return
  }
}

func main(){
  http.HandleFunc("/", MainPage)
  port := ":9999"
  error := http.ListenAndServe(":9999", nil)
  if error != nil {
    log.Fatal("ListenAndServe", error)
  }
  println("Server listen on port", port)
}
