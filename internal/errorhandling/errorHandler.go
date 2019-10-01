package errorhandling

import "fmt"

func Divide(a, b int) (float32, *someErrorType) {

	if b < 1 {
		return -1, &someErrorType{
			Time:     "time",
			Message:  fmt.Sprintf("%v cannot be less than 1", b),
			Location: 9,
		}
	}

	res := float32(float32(a)/ float32(b))
	
	return res, nil
}
