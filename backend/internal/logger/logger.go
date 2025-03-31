package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var path = fmt.Sprintf("./log/%v.log", time.Now().Format("2006-01-02_15-04-05"))

var file *os.File

func Setup() {
	var err error
	file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicln("Error opening log file:", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags)
}

func Cleanup() {
	if err := file.Close(); err != nil {
		log.Panicln("Error closing log file:", err)
	}
}
