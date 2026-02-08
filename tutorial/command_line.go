package main

import (
    "fmt"
    "os"
    // "strconv"
    "strings"
)

// func main(){
    // var s, sep string
    // for i := 1; i < len(os.Args); i++ {
	// s += sep + os.Args[i]
    	// sep = " "
    // }
    // fmt.Println(s)
    // s, sep := "", ""
    // for _, arg := range os.Args[1:]{
	// s += sep + arg
	// sep = " "
    // }

    // Ex. 1.1
    // fmt.Println(os.Args[0])
    // fmt.Println(s)
//}

// Ex 1.2
// func main(){
    // var s, sep string
    // for i := 1; i < len(os.Args); i++ {
	// s += sep + "Index: " +  strconv.Itoa(i) + " Value: " + os.Args[i]
	// sep = "\n"
    // }

    // fmt.Println(s)
// }

// Ex 1.3
func main(){
    fmt.Println(strings.Join(os.Args[1:], " "))
}
