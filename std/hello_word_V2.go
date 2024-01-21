package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Please visit http://127.0.0.1:12345/")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		s := fmt.Sprintf("你好、世界, Time is : %s", time.Now())
		_, err := fmt.Fprintf(writer, "%v\n", s)
		if err != nil {
			log.Fatal("Write Error ", err)
			return
		}
		log.Printf("%v\n", s)
	})
	if err := http.ListenAndServe("0.0.0.0:12345", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
