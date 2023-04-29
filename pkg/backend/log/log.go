package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var fileToLog = "ThesorTeX.log"

var runningInCloud = false

/**
1: ERROR
2: WARN
3: INFO
4: DEBUG
*/
var logLevel = 3

func Setup(cloud bool, level string) {
	// rm the standard timestamp from the logs
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	runningInCloud = cloud
	switch level {
	case "ERROR":
		logLevel = 1
		break
	case "WARN":
		logLevel = 2
		break
	case "DEBUG":
		logLevel = 4
		break
	default:
		logLevel = 3
	}
}

func Info(msg string, a ...any) {
	if logLevel < 3 {
		return
	}
	if a != nil {
		log.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
	} else {
		log.Println(fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), msg))
	}
}

func Error(msg string, a ...any) {
	logMessage := fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...))
	log.Println(logMessage)

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

func Warn(msg string, a ...any) {
	if logLevel < 2 {
		return
	}
	log.Println(fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
}

func Debug(msg string, a ...any) {
	if logLevel < 4 {
		return
	}
	log.Println(fmt.Sprintf("DEBUG [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
}
