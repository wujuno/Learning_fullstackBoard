package main

import (
	"fmt"
	"fullstackBoard/pkg/httpserver"
)

func main() {
	//DB연결

	//서버 시작
	fmt.Println("8080 포트에서 서버를 시작합니다🚀")
	httpserver.StartHTTPServer()
}