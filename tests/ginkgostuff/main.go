package ginkgostuff

import "fmt"

//ginkgo for BDD testing
//gomega for matchers
func main() {

}

func Sum(a, b int) (int, error) {
	return a + b, nil
}

func SumWithError(a, b int) (int, error) {

	return -1, &customError{"Overridden custom error"}
}

type customError struct {
	message string
}

func (c *customError) Error() string {
	return fmt.Sprintf(c.message)
}

