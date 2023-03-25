package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var fileToLog = "ThesorTeX.log"

var runningInCloud = false

func Setup(cloud bool) {
	runningInCloud = cloud
}

func Info(msg string, a ...any) {
	if a != nil {
		fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
	} else {
		fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), msg))
	}
}

func Error(msg string, a ...any) {
	logMessage := fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...))
	fmt.Println(logMessage)

	if !runningInCloud {
		f, err := os.OpenFile(fileToLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//defer to close when you're done with it, not because you think it's idiomatic!
		defer f.Close()

		//set output of logs to f
		log.SetOutput(f)
		log.Println(logMessage)
	}

}

func Fatal(msg string, a ...any) {
	if a != nil {
		log.Fatal(fmt.Sprintf("FATAL [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
	} else {
		log.Fatalf(fmt.Sprintf("FATAL [%v] %v", time.Now().Format("2006-01-02 15:04"), msg))
	}
}

func Debug(msg string, a ...any) {
	fmt.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
}
