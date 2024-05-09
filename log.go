package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
creates log template for us.
can pass in nil if you don't feel like typing it out,
but please log all the information as its important.
*/
func helperTemplate(header string, file any, function any, msg any, err error) string {
	return fmt.Sprintf("\n|=====%v=====|\n| File: %v\n| Function: %v\n| Msg: %v\n| Error: %v\n\n",
		header, file, function, msg, err)
}

/*
Gets filepath that it is going to log to
CANNOT USE built LOG functions as it will cause infinite recursion
*/
func getFileLocation() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%s/logs/%d-%02d-%02d.log", wd, year, int(month), day), nil
}

/*
writes log to the file. If it fails it won't crash the program.
CANNOT USE built LOG functions as it will cause infinite recursion
*/
func writeLogToFile(msg string) {
	filelocation, err := getFileLocation()
	if err != nil {
		log.Printf("\nMsg: Failed Getting File location\nError:%v\n", err)
		return
	}

	file, err := os.OpenFile(filelocation, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("\nMsg: Failed Opening File\nError:%v\n", err)
		return
	}

	defer file.Close()

	//appending time to log message as it doesn't write normally
	msgLog := fmt.Sprintf("%v%v", time.Now().Format("2006-01-02 15:04:05"), msg)

	_, err = file.WriteString(msgLog)
	if err != nil {
		log.Printf("\nMsg: Failed to write to file\nError:%v\n", err)
		return
	}
}

/*
Log a debug message without files or function params. Flag -> true message will print else it won't
*/
func DebugLog(msg string, flag bool) {
	if flag {
		logMsg := helperTemplate("DEBUG", nil, nil, msg, nil)
		writeLogToFile(logMsg)
		log.Println(logMsg)
	}
}

/*
Regular log function that takes file name and function name and user message.
Please use this over regular log so we can have this regulated.
*/
func LogNonFatal(file string, function string, msg string) {
	logMsg := helperTemplate("DEFAULT LOG", file, function, msg, nil)
	writeLogToFile(logMsg)
	log.Println(logMsg)
}

/*
Will Kill program if called.
Use when there is no direct error and you just need a message.
*/
func LogFatal(file string, function string, msg string) {
	logMsg := helperTemplate("FATAL LOG", file, function, msg, nil)
	writeLogToFile(logMsg)
	log.Fatalln(logMsg)
}

/*
We can call for every error and it will automate the background.
Does not Kill program when called.
Takes file name, function name, user message, error
*/
func LogNonFatalError(file string, function string, msg string, err error) bool {
	if err != nil {
		logMsg := helperTemplate("NON FATAL ERROR", file, function, msg, err)
		writeLogToFile(logMsg)
		log.Println(logMsg)
		return true
	}
	return false
}

/*
We can call for every error and it will automate the background.
FATAL: Kills Program if called!
Takes file name, function name, user message, error
*/
func LogFatalError(file string, function string, msg string, err error) {
	if err != nil {
		logMsg := helperTemplate("FATAL ERROR", file, function, msg, err)
		writeLogToFile(logMsg)
		log.Fatalln(logMsg)
	}
}

/*
Logs panic error nothing special just formats it
*/
func PanicError(err error) {
	if err != nil {
		logMsg := helperTemplate("PANIC ERROR", nil, nil, nil, err)
		writeLogToFile(logMsg)
		log.Panicln(logMsg)
	}
}
