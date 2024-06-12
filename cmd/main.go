package main

import (
	"fmt"
	"gra/internal"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("index.html")
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/getSettings", internal.HandleGetSettings)
	http.HandleFunc("/getStateInstance", internal.HandleGetStateInstance)
	http.HandleFunc("/sendMessage", internal.HandleSendMessage)
	http.HandleFunc("/sendFileByUrl", internal.HandleSendFileByUrl)
	fmt.Println("Server run on http://localhost:8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
