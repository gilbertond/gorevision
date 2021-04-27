package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	option1()
	option2()
	option3()
}

func option1()  {
	start:=time.Now()
	b:=new(big.Int)
	fmt.Println(b.Div(big.NewInt(1000), big.NewInt(3)))
	elapsed:=time.Since(start)
	fmt.Println(elapsed)
}

func option2()  {  // if reusable
	defer timeTrack(time.Now(), "someFunction")
	b:=new(big.Int)
	time.Sleep(100 * time.Millisecond)
	fmt.Println(b.Div(big.NewInt(1000), big.NewInt(3)))
}

func option3() { // if not needed elsewhere
	defer func(start time.Time) {
		elapsed := time.Since(start).Seconds()
		fmt.Printf("function option3 took %.fs", elapsed)
	}(time.Now())
	time.Sleep(time.Second)
}
