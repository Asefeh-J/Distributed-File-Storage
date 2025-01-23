package main

import (
	"fmt"
	"os"

	"github.com/Asefeh-J/Distributed-File-Storage/shared/logger"
)

func InitLogger() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("can't get working directory: %v", err)
		os.Exit(-1)
	}
	logger.InitLog(path, "file-service.log")
	logger.Inst().Info("file-service logger initialized")

}

func Init() {
	InitLogger()
}

func main() {
	Init()
}
