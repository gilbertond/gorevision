package errorhandling

import (
	"fmt"
)

type someErrorType struct {
	Message  string
	Time     string
	Location int
}

	func (e *someErrorType) Error() string {
	return fmt.Sprintf("Time: %s, Location %d, Error %s ", e.Time, e.Location, e.Message)
}
