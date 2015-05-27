package main

import (
	"fmt"

	"github.com/immesys/bw2bind"
)

func main() {
	_, err := bw2bind.Connect("127.0.0.1:28589")
	if err != nil {
		fmt.Println("Got error:", err)
	}

	//Create an entity for ourselves

}
