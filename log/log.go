package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogPath struct {
	file     string
	function string
}

var (
	_filePath string
	logPaths  []LogPath
)

/*
Sets a new log with file and function.
Returns the code which will be used to refernce the file and function name.
*/
func SetLog(file, function string) int {
	code := len(logPaths)

	logPaths = append(logPaths, LogPath{
		file:     file,
		function: function,
	})

	return code
}

/*
Returns true if code should exist.
Returns False if code should not exist.
*/
func validateCode(code int) bool {
	// checks if its invalid and flips result
	return !(code < 0 || code >= len(logPaths)) // if not invalid
}

/*
Uses the code to print file and function.
Can be sussy if not used properly and by passing in random codes.
Only use when you have SetLog().
*/
func LogNonFatal_C(code int, msg string) {
	if !validateCode(code) {
		LogNonFatal("log.go", "LogNonFatal_C()", "Entered Invalid Code. Returning False automacially")
		return
	}

	logPath := logPaths[code]
	LogNonFatal(logPath.file, logPath.function, msg)
}

/*
Uses the code to print file and function.
Can be sussy if not used properly and by passing in random codes.
Only use when you have SetLog().
*/
func LogNonFatalError_C(code int, msg string, err error) bool {
	if !validateCode(code) {
		LogNonFatal("log.go", "LogNonFatalError_C()", "Entered Invalid Code. Returning False automacially")
		return false
	}

	logPath := logPaths[code]
	return LogNonFatalError(logPath.file, logPath.function, msg, err)
}

/*
pass in the get working directory argument and it will find the logs folder from there.
os.Getwd()
*/
func SetFilePath(dir string) {
	_filePath = dir
}

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
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%s/logs/%d-%02d-%02d.log", _filePath, year, int(month), day), nil
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
