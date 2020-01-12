package main

import (
	"fmt"
	"time"
)

const loop_count = 100000

func showHelp() {
	fmt.Printf(
		`
    ============================================================================================================
    This program measures difference of processing-time between passing-by-value and passing-by-pointer methods.
    Both methods do the following equivalent operation.
    - use a struct which has [%d][%d]int array
    - incerement all element in the array %d times

    Passing-by-Value version returns new array, on the other hand,
    passing-by-pointer version directly changes a struct which is given as input argument.
    ============================================================================================================`, size, size, loop_count)
}

func main() {
	showHelp()
	fmt.Println("")

	var start, end time.Time
	var tmp TestStruct

	/* Passing-by-Value Version */
	start = time.Now()
	for i := 0; i < loop_count; i++ {
		tmp = tmp.IncrementByValue()
	}
	end = time.Now()

	fmt.Printf("Elapsed time(passing-by-value): %f [sec]\n",
		end.Sub(start).Seconds())
	//tmp.PrintData()

	/* Passing-by-Pointer Version*/
	start = time.Now()
	for i := 0; i < loop_count; i++ {
		tmp.IncrementByPointer()
	}
	end = time.Now()
	fmt.Printf("Elapsed time(passing-by-pointer): %f [sec]\n",
		end.Sub(start).Seconds())
	//tmp.PrintData()
}
