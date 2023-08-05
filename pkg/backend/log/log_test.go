package log

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestLogLevel_INFO(t *testing.T) {
	Setup(true, "INFO")

	for _, tc := range []struct {
		level LogLevel
		exp   string
	}{
		{
			level: ERROR,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: WARN,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: INFO,
			exp:   fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: DEBUG,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			log.SetOutput(&str)

			switch tc.level {
			case ERROR:
				Error("hello")
				break
			case WARN:
				Warn("hello")
				break
			case INFO:
				Info("hello")
				break
			case DEBUG:
				Debug("hello")
				break
			}

			got := strings.TrimSuffix(str.String(), "\n")

			if got != tc.exp {
				t.Errorf("expected %s, got %s", tc.exp, got)
			}
		})
	}

}

func TestLogLevel_DEBUG(t *testing.T) {
	Setup(true, "DEBUG")

	for _, tc := range []struct {
		level LogLevel
		exp   string
	}{
		{
			level: ERROR,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: WARN,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: INFO,
			exp:   fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: DEBUG,
			exp:   fmt.Sprintf("DEBUG [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			log.SetOutput(&str)

			switch tc.level {
			case ERROR:
				Error("hello")
				break
			case WARN:
				Warn("hello")
				break
			case INFO:
				Info("hello")
				break
			case DEBUG:
				Debug("hello")
				break
			}

			got := strings.TrimSuffix(str.String(), "\n")

			if got != tc.exp {
				t.Errorf("expected %s, got %s", tc.exp, got)
			}
		})
	}

}

func TestLogLevel_WARN(t *testing.T) {
	Setup(true, "WARN")

	for _, tc := range []struct {
		level LogLevel
		exp   string
	}{
		{
			level: ERROR,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: WARN,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: INFO,
			exp:   "",
		},
		{
			level: DEBUG,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			log.SetOutput(&str)

			switch tc.level {
			case ERROR:
				Error("hello")
				break
			case WARN:
				Warn("hello")
				break
			case INFO:
				Info("hello")
				break
			case DEBUG:
				Debug("hello")
				break
			}

			got := strings.TrimSuffix(str.String(), "\n")

			if got != tc.exp {
				t.Errorf("expected %s, got %s", tc.exp, got)
			}
		})
	}

}

func TestLogLevel_ERROR(t *testing.T) {
	Setup(true, "ERROR")

	for _, tc := range []struct {
		level LogLevel
		exp   string
	}{
		{
			level: ERROR,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: WARN,
			exp:   "",
		},
		{
			level: INFO,
			exp:   "",
		},
		{
			level: DEBUG,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			log.SetOutput(&str)

			switch tc.level {
			case ERROR:
				Error("hello")
				break
			case WARN:
				Warn("hello")
				break
			case INFO:
				Info("hello")
				break
			case DEBUG:
				Debug("hello")
				break
			}

			got := strings.TrimSuffix(str.String(), "\n")

			if got != tc.exp {
				t.Errorf("expected %s, got %s", tc.exp, got)
			}
		})
	}

}
