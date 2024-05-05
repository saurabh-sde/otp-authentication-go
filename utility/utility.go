package utility

import (
	"log"
)

func Print(err *error, v ...interface{}) {
	if len(v) > 0 {
		// print INFO
		log.Printf("INFO: %+v\n", v)
	}
	// print error
	if err != nil {
		log.Printf("ERROR: %+v\n", *err)
		return
	}
}
