package log

import (
	"encoding/json"
	simpleerror "main/Log/SimpleErorr"
	richerror "main/RichErorr"
	"os"

	"time"
)

type Log struct {
	Errors []richerror.RichError
}

func (l *Log) Append(err error) {
	var finalError richerror.RichError

	// type assertion
	rErr, ok := err.(*richerror.RichError)
	if ok {
		finalError = *rErr
	} else {
		sErr, ok := err.(*simpleerror.SimpleError)
		if ok {
			finalError = richerror.RichError{
				Message:   sErr.Output,
				MetaData:  nil,
				Operation: sErr.Operation,
				Time:      time.Now(),
			}
		} else {
			finalError = richerror.RichError{
				Message:   err.Error(),
				MetaData:  nil,
				Operation: "unknown",
				Time:      time.Now(),
			}
		}

	}

	l.Errors = append(l.Errors, finalError)
}

func (l *Log) Save() {
	//for i, e := range l.Errors {
	//	fmt.Printf("i: %d, operation: %s, message: %s, meta-data: %+v\n",
	//		i, e.Operation, e.Message, e.MetaData)
	//}

	f, _ := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()

	data, _ := json.Marshal(l.Errors)
	f.Write(data)
}
