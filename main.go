package main

import (
	"gocore/framework"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("go frame 启动失败")
	}
}
