package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Fatal *log.Logger
)

func Init() {
	Info = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
	Fatal = log.New(os.Stderr, "FATAL:", log.Ldate|log.Ltime|log.Lshortfile)
}
