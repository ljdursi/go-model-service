/*
Package errors implements error reporting functionality that is commonly used within this project.
Its purpose is to standardize generic error reports, both for external use via API responses and for
internal use via logging.
 */
package errors

import "log"

// Log employs the log package to print an error message in a consistent format.
func Log(err error, httpCode int32, funcName string, message string) {
	log.Printf("%d ERROR: %s \n" +
		"IN: %s \n",
		httpCode, message, funcName)
	if err != nil {
		log.Println("ERROR MESSAGE FOLLOWS:\n" + err.Error())
	}
}