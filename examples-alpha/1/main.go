package main

import (
	"fmt"
	"os"
	"time"

	bw "gopkg.in/immesys/bw2bind.v2"
)

func main() {
	// This will connect to local router on localhost:28589
	cl := bw.ConnectOrExit("")

	// This will set the entity file for this connection. You can only
	// have one entity per connection. If you have multiple identities
	// (like a splicer) then you need multiple clients. This is because
	// the SSL connection verification is tied to your identity
	vk := cl.SetEntityFileOrExit("ex1.key")
	// vk is the string form of our Verifying Key
	fmt.Println("entity set: ", vk)

	// We will discuss this in example 2
	cl.OverrideAutoChainTo(true)

	// The key in this example is allowed to use castle.bw2.io/bwtools/*
	// Lets subscribe to everything on that uri
	msgchan, err := cl.Subscribe(&bw.SubscribeParams{URI: "oski.demo/bwtools/*"})
	if err != nil {
		// This error could be things like you not having permissions
		// or not being able to connect to the remote router etc
		fmt.Println("Failed to subscribe: ", err)
		os.Exit(1)
	}

	// Now lets consume and print messages. All calls are blocking,
	// so use a goroutine
	go func() {
		for m := range msgchan {
			// Dump() is a utility method that determines the type of all
			// the POs in the message and prints them out as best it can
			m.Dump()
		}
	}()

	// To give some messages to see (although you may see other
	// people's) lets also publish some messages
	// First lets make a message. Remember that all bosswave messages
	// are sequences of strongly typed Payload Objects. Lets do
	// a string and a messagepack object. For details on the available
	// PO types (and to create more) look at github.com/immesys/bw2_pid
	po1 := bw.CreateStringPayloadObject("Hello world!")
	// LogDict is a sub-type of MsgPack with a little more semantic meaning
	po2, _ := bw.CreateMsgPackPayloadObject(bw.PONumLogDict, map[string]interface{}{
		"example": 1,
		"hello":   "world",
	})
	for {
		err = cl.Publish(&bw.PublishParams{
			URI:            "oski.demo/bwtools/1",
			PayloadObjects: []bw.PayloadObject{po1, po2},
		})
		fmt.Println("Published, err was: ", err)
		time.Sleep(5 * time.Second)
	}

}
