//Echo3 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main(){
	for key, arg := range os.Args[1:]{
		fmt.Println(key, arg)
	}
}

