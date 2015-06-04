package main

import (
	"fmt"
	"os"

	"github.com/immesys/bw2bind"
)

func bail(v interface{}) {
	fmt.Println(v)
	os.Exit(1)
}
func main() {
	cl, err := bw2bind.Connect("localhost:28589")
	if err != nil {
		bail(err)
	}
	err = cl.SetEntityFile("eve.key")
	if err != nil {
		bail(err)
	}
	msgchan, err := cl.Subscribe(&bw2bind.SubscribeParams{
		URI:                "castle.bw2.io/ex/baz",
		PrimaryAccessChain: "a6AYZos9UZClIJ-rD5Azrmij64qbT0q8xbGSRmhqKrU=",
	})
	fmt.Println("subscribe finished")
	if err != nil {
		bail(err)
	}
	for m := range msgchan {
		m.Dump()
	}
	fmt.Println("channel ended")
}
