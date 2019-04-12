package Base

import (
	"SORA/Config"
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	DBTime  *log.Logger
)

func LogInit(returnerr chan error) {
	os.MkdirAll("log", os.ModePerm)
	// -----------------------------------------------[Error log]
	if Config.DebugLevel >= 1 {
	} else if errSave, err := os.OpenFile("./log/Error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalf("error opening file: %v", err)
		returnerr <- err
	} else {
		Error = log.New(io.MultiWriter(os.Stderr, errSave), "【Error】 ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	// -----------------------------------------------[Warning Log]
	if Config.DebugLevel >= 2 {
	} else if warningSave, err := os.OpenFile("./log/Warning.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalf("error opening file: %v", err)
		returnerr <- err
	} else {
		Warning = log.New(io.MultiWriter(os.Stderr, warningSave), "【Warning】 ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	// -----------------------------------------------[Server Log]
	if Config.DebugLevel >= 3 {
	} else if infoSave, err := os.OpenFile("./log/Info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalf("error opening file: %v", err)
		returnerr <- err
	} else {
		Info = log.New(io.MultiWriter(os.Stderr, infoSave), "【Info】 ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	// -----------------------------------------------[Database time Log]
	if Config.DebugLevel >= 4 {
	} else if dbInfo, err := os.OpenFile("./log/Database_Time.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalf("error opening file: %v", err)
		returnerr <- err
	} else {
		DBTime = log.New(io.MultiWriter(dbInfo), "【DB Time】 ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
