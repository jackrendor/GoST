package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	// DATEFORMAT Date format when printing/writing logs
	DATEFORMAT string = "2006-01-02 15:04:05"
	// USAGESTR Usage message
	USAGESTR string = "Usage: %s <url> <request delay (seconds)> <logfile (optional)>\n"
)

func exitOnArgLack(args []string) {
	// Make the program exit if argument are not enough
	if len(args) < 3 {
		fmt.Printf(USAGESTR, args[0])
		os.Exit(-1)
	}
}
func sleep(sec int) {
	time.Sleep(time.Second * time.Duration(sec))
}

func writeLog(line string, logFileName string) {
	// Access type set on "append" by default
	mode := (os.O_APPEND | os.O_WRONLY)

	// Check if file exist.
	_, err := os.Stat(logFileName)
	if os.IsNotExist(err) {
		// If file doesn't exist, set mode type to "write"
		mode = (os.O_CREATE | os.O_WRONLY)
	}
	// Open file.
	f, err := os.OpenFile(logFileName, mode, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Try to write data on file. If fails, exit.
	if _, err = f.WriteString(line); err != nil {
		panic(err)
	}
}
func printLog(statusCode int, notes string, logFileName string) {
	// Get the current time
	now := time.Now().Format(DATEFORMAT)

	// Format the string in order to log stuff as we like
	var line string
	if statusCode >= 200 && statusCode <= 299 {
		line = fmt.Sprintf("%s SUCCESS %d\n", now, statusCode)
	} else if statusCode >= 400 && statusCode <= 499 {
		line = fmt.Sprintf("%s NOTFOUND %d\n", now, statusCode)
	} else if statusCode >= 500 && statusCode <= 599 {
		line = fmt.Sprintf("%s SRV_ERROR %d\n", now, statusCode)
	} else {
		line = fmt.Sprintf("%s UNKOWN %d %s\n", now, statusCode, notes)
	}
	// Print the log
	fmt.Printf("%s", line)

	// If we specified the filename, write logs in a file
	if logFileName != "" {
		writeLog(line, logFileName)
	}
}

func makeReq(url string, client *http.Client, logFileName string) {
	resp, netErr := client.Get(url)
	if netErr != nil {
		printLog(-1, netErr.Error(), logFileName)
		return
	}
	defer resp.Body.Close()

	printLog(resp.StatusCode, "", logFileName)

}

func main() {
	exitOnArgLack(os.Args)

	// Get argument passed
	target := os.Args[1]
	delay, delayErr := strconv.Atoi(os.Args[2])
	logFileName := ""

	// If arguments are more than 3, user specified logfile.
	if len(os.Args) > 3 {
		logFileName = os.Args[3]
	}

	fmt.Println("Starting routine...")

	// If argument delay isn't a number, exit.
	if delayErr != nil {
		fmt.Println("Please, second argument must be integer.")
		os.Exit(-1)
	}

	// A little edit to the client to ignore SSL errors.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	for true {
		makeReq(target, client, logFileName)
		sleep(delay)
	}

}
