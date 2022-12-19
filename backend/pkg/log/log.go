package log

import (
	"fmt"
	"time"
)

func Info(msg string, a ...any) {
	if a != nil {
		fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
	} else {
		fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), msg))
	}
}

func Error(msg string, a ...any) {
	fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
}

func Debug(msg string, a ...any) {
	fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
}
