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
	fmt.Println("connected")
	us, err := cl.SetEntityFile("eve.key")
	if err != nil {
		bail(err)
	}
	fmt.Println("entity set: ", us)
	uri := "michael.bw2.io/ex/5/*"

	//Build a chain
	rc, err := cl.BuildChain(uri, "C*", us)
	if err != nil {
		bail(err)
	}
	fmt.Println("BC returned")
	pac := ""
	for chain := range rc {
		fmt.Printf("Got chain: %v", chain.Hash)
		pac = chain.Hash
	}
	fmt.Println("chain build rv")
	if pac == "" {
		bail("Could not find a PAC")
	}
	//Subscribe
	msgchan, err := cl.Subscribe(&bw2bind.SubscribeParams{
		URI:                uri,
		PrimaryAccessChain: pac,
	})
	if err != nil {
		bail(err)
	}

	fmt.Println("subscribe finished")
	if err != nil {
		bail(err)
	}
	for m := range msgchan {
		m.Dump()
	}
	fmt.Println("channel ended")
}
