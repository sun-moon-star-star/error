package errors

import "error"

type API struct {
	RecordError ([]byte) error // Record error
	Retry(*caller.Caller) error // 
}
