package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	target := "here is HYQWN Spqce"
	fmt.Fprintf(w, "Hello, %s!\n", target)
}

func main() {
	log.Print("微信云托管服务启动成功")
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
