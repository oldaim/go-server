package main

import (
	"fmt"
	"go-server/database"
	"log"
)

func main() {

	if err := database.Initialize(); err != nil {
		log.Fatal("데이터베이스 연결 실패:", err)
	}

	fmt.Println("Hello, World!")

}
