package log

import (
	"fmt"
	"time"
)

func Info(msg ...interface{}) {
	fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprint(msg...)))
}

func Error(msg ...interface{}) {
	fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprint(msg...)))
}

func Debug(msg ...interface{}) {
	fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprint(msg...)))
}
