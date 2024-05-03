package utility

import "fmt"

func Print(err *error, v ...interface{}) {
	if len(v) > 0 {
		// print INFO
		fmt.Printf("INFO: %+v\n", v)
	}
	// print error
	if err != nil {
		fmt.Printf("ERROR: %+v\n", *err)
		return
	}
}
