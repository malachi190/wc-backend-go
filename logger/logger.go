package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Error *log.Logger
	Info  *log.Logger
	Debug *log.Logger

	Console *log.Logger

	once sync.Once
)

// Init sets up a daily-rotated error log using lumberjack.
func Init(dir string) {
	once.Do(func() { setup(dir) })
}

func setup(dir string) {
	_ = os.MkdirAll(dir, 0755)

	today := time.Now().Format("2006-01-02")

	filename := filepath.Join(dir, fmt.Sprintf("error-%s.log", today))

	fileWriter := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    500, // MB
		MaxBackups: 30,  // keep last 30 daily files
		LocalTime:  true,
		Compress:   true,
	}

	flag := log.LstdFlags | log.Lshortfile

	Error = log.New(fileWriter, "[ERROR] ", flag)
	Info = log.New(fileWriter, "[INFO] ", flag)
	Debug = log.New(fileWriter, "[DEBUG] ", flag)

	Console = log.New(os.Stdout, "[REQUEST] ", log.LstdFlags)
}
