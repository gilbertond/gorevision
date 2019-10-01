package main

import (
	"fmt"
	"github.com/gorevision/internal/errorhandling"
	"github.com/pkg/errors"
	"log"
)

func main() {

	res, error := errorhandling.Divide(6, 0)
	//if error != nil {
	//	log.Fatal(error)
	//}

	if error != nil {
		log.Fatal(

			//good way to return error to calling client
			errors.Wrap(error, "something failed"),
		)
	}

	fmt.Printf("Divide 6/4=%v \n", res)
	fmt.Println()
	fmt.Printf("Hello world\n")

	var fruit [3] string

	fruit[0] = "1233"
	fruit[1] = "1233"
	fruit[2] = "1233"

	//Declare and assign/initialize
	fruitArr := []string{"1", "2", "3", "4"}

	fmt.Println(fruitArr)

	var x, y = "am x", "am y"

	fmt.Println(x + " and " + y)
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[1:3]) //Range

	//Switch
	switch y {
	case "x":
		fmt.Println("x")
	case "y":
		fmt.Println("y")
	default:
		fmt.Println("default")
	}

	//for loop
	for i := 0; i < len(fruitArr); i++ {
		fmt.Println(fruitArr[i])
	}

	//Maps
	mapData := make(map[string]string)
	mapData2 := map[string]string{
		"key1": "value1", "key2": "value2"} //Breaking this line forces you to add a comma. funny Go...

	mapData["key1"] = "sds"
	mapData["key2"] = "dsfd"
	mapData["key3"] = "fdsas"
	fmt.Println(mapData2["key1"])

	delete(mapData2, "key1")
	mapData2["key33"] = "data33"

	//loop array/map and print index:value
	for i, id := range mapData { // i and id in this case have to be used
		fmt.Printf("%s %s \n", i, id)
	}

	for key := range mapData {
		fmt.Printf("%s\n", key)
	}

	for _, value := range mapData {
		fmt.Printf("%s\n", value)
	}

	ids := []int{33, 76, 54, 23, 11, 2}
	for i, id := range ids { // i and id in this case have to be used
		fmt.Printf("%d %d \n", i, id)
	}

	sum := 0
	for _, id := range ids { //The underscores solves issue with above
		fmt.Printf("%d \n", id)
		sum += id
	}

	//	Pointers
	a := 2
	b := &a //assigns memory address of a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) //%T for type
	fmt.Println(*b)       //print value of memory address; 2

	*b = 10        //change memory address value
	fmt.Println(a) // value of a is now 10

	//	Closure function
	sum2 := clouserFunction()
	for i := 0; i < 10; i++ {
		fmt.Println(sum2(i))
	}
}

//	Closures: functions return function
func clouserFunction() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x

		return sum
	}
}
