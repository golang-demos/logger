package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "[INFO]   : ", log.LstdFlags|log.Lshortfile)
	WarningLogger = log.New(file, "[WARNING]: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "[ERROR]  : ", log.LstdFlags|log.Lshortfile)
}

func main() {
	// 2022/02/27 02:28:34 This default log message
	log.Println("This is a default log message")

	// log.go:9: This is log message with short path of file
	log.SetFlags(log.Lshortfile)
	log.Println("This is a log message with short path of file")

	// 05:08:47.731696 This is log message with micro seconds
	log.SetFlags(log.Lmicroseconds)
	log.Println("This is a log message with micro seconds")

	// 05:09:18.310000 log.go:9: This my log message
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	log.Println("This is a log message with both microseconds and short file path")

	// 05:15:38.266747 log.go:23: Error Occured
	// exit status 1
	// log.Fatal("Error Occured")

	// 05:15:07.906567 log.go:25: Panic raised
	// panic: Panic raised
	// log.Panic("Panic raised")

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("This message is written to a file.")

	InfoLogger.Println("This is a INFO Log message")
	WarningLogger.Println("This is a WARNING Log message")
	ErrorLogger.Println("This is a ERROR Log message")

}
