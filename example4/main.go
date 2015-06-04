package main

import (
	"bufio"
	"fmt"
	"os"

	bw "github.com/immesys/bw2bind"
)

func bail(v interface{}) {
	fmt.Println(v)
	os.Exit(1)
}
func main() {
	cl, err := bw.Connect("localhost:28589")
	if err != nil {
		bail(err)
	}
	err = cl.SetEntityFile("ex2.key")
	if err != nil {
		bail(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		err = cl.Publish(&bw.PublishParams{
			URI:                "castle.bw2.io/ex/baz",
			PrimaryAccessChain: "fVpZ4fEJZuPYEjUTOXqg7Xon9rvzss6ZyeyE-ZsyJZ0=",
			PayloadObjects:     []bw.PayloadObject{bw.CreateStringPayloadObject(text)},
		})
		if err != nil {
			bail(err)
		}
		fmt.Printf("Published ok")
	}
}
