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
	us, err := cl.SetEntityFile("ex2.key")
	if err != nil {
		bail(err)
	}
	uri := "castle.bw2.io/ex/baz"
	pac, _ := cl.BuildAnyChain(uri, "P", us)
	if pac == nil {
		fmt.Println("Could not get permissions")
		os.Exit(1)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		err = cl.Publish(&bw.PublishParams{
			URI:                uri,
			PrimaryAccessChain: pac.Hash,
			DoVerify:           true,
			PayloadObjects:     []bw.PayloadObject{bw.CreateStringPayloadObject(text)},
		})
		if err != nil {
			bail(err)
		}
		fmt.Printf("Published ok")
	}
}
