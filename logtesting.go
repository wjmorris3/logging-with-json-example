package main

import (
	"fmt"
	"net/http"
	"os"

	uuid "github.com/google/uuid"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {

	/*
		This generates a random UID that will be passed with each part of the transaction, ensuring complete tracing.
	*/
	transactionKey, err := uuid.NewRandom()
	fmt.Printf("Random UUID Transaction Key: %s\n", transactionKey)

	//error handling required to generate the initial UID (transactionKey)
	if err != nil {
	}

	//setup for printing map objects
	loggingTestRequest, _ := http.NewRequest(http.MethodGet, "/test/webapp", nil)
	loggingTestRequest.Header.Add("user-agent", "some-browser")
	loggingTestRequest.Header.Add("user-agent", "another-browser")
	loggingTestRequest.Header.Add("x-forwarded-ip", "127.0.0.1")
	loggingTestRequest.Header.Add("x-forwarded-ip", "192.168.1.1")
	loggingTestRequest.Header.Add("some-data", "http-data")

	/*
		START demo for context logging
		this is a general message that can be re-used by using the contextLogger. This should be used if you want the same message logged for multiple log levels
		FATAL and PANIC log levels will cause the application to quit and should not be used unless the intention is to end the app

		Sample output for INFO log level:
		{
		"level":"info",
		"msg":"info log message",
		"sample message":"sample logging message using Context Logging",
		"sso":"444444444",
		"time":"2018-11-26T12:00:04-05:00",
		"uid":"5cf8ed31-b894-43ab-82d6-35c97accb312"
		}
	*/

	contextLogger := log.WithFields(log.Fields{
		"logMapObject":   loggingTestRequest.Header,
		"sso":            "123123123",
		"sample message": "sample logging message using Context Logging",
		"uid":            transactionKey,
	})

	fmt.Println("*** Begin Context logging example ***")
	fmt.Println()

	contextLogger.Info("info log message")
	fmt.Println()

	contextLogger.Debug("info log message")
	fmt.Println()

	contextLogger.Warn("warn log message")
	fmt.Println()

	contextLogger.Error("error log message")
	fmt.Println()

	contextLogger.Trace("trace log message")
	fmt.Println()

	/*
		Enable for Fatal and Panic logging levels

		contextLogger.Fatal("fatal log message")
		fmt.Println()

		contextLogger.Panic("panic log message")
		fmt.Println()
	*/

	/*
		Begin standard logging demo
		Setup the customfomatter.
		The time stamp can be configured with this method.
		Using standard Unix time format.
	*/
	log.SetFormatter(&log.JSONFormatter{})
	customFormatter := new(logrus.JSONFormatter)
	customFormatter.TimestampFormat = "Mon Jan _2 15:04:05 MST 2006"

	log.SetFormatter(customFormatter)

	/*
		Enable this line for logging with function/method name included in the log.
		The 'name' in the Json name value pair is sorted alphabetically and determines the order of the messages.
		Uppercase is first, then lowercase names.
		Warning this is considered a HEAVY call and adds overhead.

		log.SetReportCaller(true)
	*/
	fmt.Println("*** Begin individual log lines example ***")
	fmt.Println()

	log.WithFields(log.Fields{
		"testMapObj": loggingTestRequest.Header,
		"sso":        "123456789",
		"size":       10,
		"uid":        transactionKey,
	}).Info("Info log message")
	fmt.Println()

	log.WithFields(log.Fields{
		"sso":  "999999999",
		"size": 25,
		"uid":  transactionKey,
	}).Warn("Warn log message")
	fmt.Println()

	log.WithFields(log.Fields{
		"sso":   "555555555",
		"size":  45,
		"title": "debug title",
		"moo":   "m00",
		"uid":   transactionKey,
	}).Debug("Debug log message")
	fmt.Println()

	/*
		FATAL and Panic call exit, only use this if you want the application to end.

		log.WithFields(log.Fields{
			"sso":  "555555555",
			"size": 62,
			"type": "Fatal",
			"uid":  transactionKey,
		}).Fatal("Fatal log message")
		fmt.Println()

		log.WithFields(log.Fields{
			"sso":  "555555555",
			"size": 88,
			"type": "Fatal",
			"uid":  transactionKey,
		}).Panic("Panic log message")
		fmt.Println()
	*/
}
