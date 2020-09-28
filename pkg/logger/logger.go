package logger

import (
	"log"
	"os"
)

var (
	// INFO logger
	INFO = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
)
