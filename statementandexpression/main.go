//package main is being declared in the execute packages
package main

import (
	"fmt"
	"runtime"
)

/*
	main function is execited by the go complier automatically
*/
func main() {
	fmt.Println("No of CPU in this system:", runtime.NumCPU())
}
