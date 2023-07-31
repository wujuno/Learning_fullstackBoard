package main

import (
	"fmt"
	"fullstackBoard/db"
	"fullstackBoard/pkg/httpserver"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//DB연결
	fmt.Println("DB에 연결합니다.")
	db, err := db.Init()
	if err != nil {
		fmt.Println("DB에 연결에 실패했습니다.:", err)
		os.Exit(1)
	}
	defer db.Close()

	//서버 시작
	fmt.Println("8080 포트에서 서버를 시작합니다🚀")
	httpserver.StartHTTPServer()
}