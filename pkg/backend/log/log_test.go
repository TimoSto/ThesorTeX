package log_test

import (
	"bytes"
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	golog "log"
	"strings"
	"testing"
	"time"
)

func TestLogLevel_INFO(t *testing.T) {
	log.Setup(true, "INFO")

	for _, tc := range []struct {
		level log.LogLevel
		exp   string
	}{
		{
			level: log.ERROR,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.WARN,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.INFO,
			exp:   fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.DEBUG,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			golog.SetOutput(&str)

			switch tc.level {
			case log.ERROR:
				log.Error("hello")
				break
			case log.WARN:
				log.Warn("hello")
				break
			case log.INFO:
				log.Info("hello")
				break
			case log.DEBUG:
				log.Debug("hello")
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
	log.Setup(true, "DEBUG")

	for _, tc := range []struct {
		level log.LogLevel
		exp   string
	}{
		{
			level: log.ERROR,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.WARN,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.INFO,
			exp:   fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.DEBUG,
			exp:   fmt.Sprintf("DEBUG [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			golog.SetOutput(&str)

			switch tc.level {
			case log.ERROR:
				log.Error("hello")
				break
			case log.WARN:
				log.Warn("hello")
				break
			case log.INFO:
				log.Info("hello")
				break
			case log.DEBUG:
				log.Debug("hello")
				break
			}

			got := strings.TrimSuffix(str.String(), "\n")

			if got != tc.exp {
				t.Errorf("expected %s, got %s", tc.exp, got)
			}
		})
	}

}

func TestLogLevel_WARNING(t *testing.T) {
	log.Setup(true, "WARNING")

	for _, tc := range []struct {
		level log.LogLevel
		exp   string
	}{
		{
			level: log.ERROR,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.WARN,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.INFO,
			exp:   "",
		},
		{
			level: log.DEBUG,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			golog.SetOutput(&str)

			switch tc.level {
			case log.ERROR:
				log.Error("hello")
				break
			case log.WARN:
				log.Warn("hello")
				break
			case log.INFO:
				log.Info("hello")
				break
			case log.DEBUG:
				log.Debug("hello")
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
	log.Setup(true, "ERROR")

	for _, tc := range []struct {
		level log.LogLevel
		exp   string
	}{
		{
			level: log.ERROR,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: log.WARN,
			exp:   "",
		},
		{
			level: log.INFO,
			exp:   "",
		},
		{
			level: log.DEBUG,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			golog.SetOutput(&str)

			switch tc.level {
			case log.ERROR:
				log.Error("hello")
				break
			case log.WARN:
				log.Warn("hello")
				break
			case log.INFO:
				log.Info("hello")
				break
			case log.DEBUG:
				log.Debug("hello")
				break
			}

			got := strings.TrimSuffix(str.String(), "\n")

			if got != tc.exp {
				t.Errorf("expected %s, got %s", tc.exp, got)
			}
		})
	}

}
