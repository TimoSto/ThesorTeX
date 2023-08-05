package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var fileToLog = "ThesorTeX.log"

var runningInCloud = false

type LogLevel int

const (
	ERROR LogLevel = iota
	WARN
	INFO
	DEBUG
)

var logLevel = DEBUG

func Setup(cloud bool, level string) {
	// rm the standard timestamp from the logs
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	runningInCloud = cloud
	switch level {
	case "ERROR":
		logLevel = ERROR
		break
	case "WARNING":
		logLevel = WARN
		break
	case "INFO":
		logLevel = INFO
		break
	default:
		logLevel = DEBUG
	}

	Debug("set log level to %v", logLevel)
}

func Info(msg string, a ...any) {
	if logLevel < INFO {
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
	if logLevel < WARN {
		return
	}
	log.Println(fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
}

func Debug(msg string, a ...any) {
	if logLevel < DEBUG {
		return
	}
	log.Println(fmt.Sprintf("DEBUG [%v] %v", time.Now().Format("2006-01-02 15:04"), fmt.Sprintf(msg, a...)))
}
