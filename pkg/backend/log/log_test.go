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
		level int
		exp   string
	}{
		{
			level: 1,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 2,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 3,
			exp:   fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 4,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			log.SetOutput(&str)

			switch tc.level {
			case 1:
				Error("hello")
				break
			case 2:
				Warn("hello")
				break
			case 3:
				Info("hello")
				break
			case 4:
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
		level int
		exp   string
	}{
		{
			level: 1,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 2,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 3,
			exp:   fmt.Sprintf("INFO [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 4,
			exp:   fmt.Sprintf("DEBUG [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			log.SetOutput(&str)

			switch tc.level {
			case 1:
				Error("hello")
				break
			case 2:
				Warn("hello")
				break
			case 3:
				Info("hello")
				break
			case 4:
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
		level int
		exp   string
	}{
		{
			level: 1,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 2,
			exp:   fmt.Sprintf("WARN [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 3,
			exp:   "",
		},
		{
			level: 4,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			log.SetOutput(&str)

			switch tc.level {
			case 1:
				Error("hello")
				break
			case 2:
				Warn("hello")
				break
			case 3:
				Info("hello")
				break
			case 4:
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
		level int
		exp   string
	}{
		{
			level: 1,
			exp:   fmt.Sprintf("ERROR [%v] %v", time.Now().Format("2006-01-02 15:04"), "hello"),
		},
		{
			level: 2,
			exp:   "",
		},
		{
			level: 3,
			exp:   "",
		},
		{
			level: 4,
			exp:   "",
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.level), func(t *testing.T) {
			var str bytes.Buffer

			log.SetOutput(&str)

			switch tc.level {
			case 1:
				Error("hello")
				break
			case 2:
				Warn("hello")
				break
			case 3:
				Info("hello")
				break
			case 4:
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
