package logging

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"log"
	"os"
	"time"
)

var (
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
)

func initLoggSetting() {
	LogSavePath = setting.AppSetting.LogSavePath
	LogSaveName = setting.AppSetting.LogSaveName
	LogFileExt = setting.AppSetting.LogFileExt
	TimeFormat = "20060102"
}

func getLogFilePath() string {
	return LogSavePath
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("permission: %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to open: %v", err)
	}

	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
