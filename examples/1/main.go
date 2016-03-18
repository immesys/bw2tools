package main

import (
	"fmt"
	"time"

	bw "github.com/immesys/bw2bind"
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
	msgchan, err := cl.Subscribe(&bw.SubscribeParams{URI: "castle.bw2.io/bwtools/*"})
	if err != nil {
		// This error could be things like you not having permissions
		// or not being able to connect to the remote router etc
		panic(err)
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
			URI:            "castle.bw2.io/bwtools/1",
			PayloadObjects: []bw.PayloadObject{po1, po2},
		})
		fmt.Println("Published, err was: ", err)
		time.Sleep(5 * time.Second)
	}
	// }
	// 	err = cl.Publish(&bw.PublishParams{
	// 		URI:                uri,
	// 		PrimaryAccessChain: pac.Hash,
	// 		DoVerify:           true,
	// 		ElaboratePAC:       bw.ElaborateFull,
	// 		Persist:            true,
	// 		PayloadObjects:     []bw.PayloadObject{bw.CreateStringPayloadObject("Hello world")},
	// 	})
	// 	if err != nil {
	// 		fmt.Println("finished persisting err: ", err.Error())
	// 	} else {
	// 		fmt.Println("Finished persisting, no error")
	// 	}
	//
	//
	//   // AutoChain asks our local router to build the required
	//   // chain of trust for our action. In example 2 we will look
	//   // at alternatives
	//
	//
	//    It makes sense for most
	//   // subscribe operations or other one-offs. If you are going
	//   // to be doing multiple operations using the same credentials
	//   // such as lots of publishes, then you should rather build
	//   // a chain by itself and cache the result. For example:
	//
	//   // You can build multiple chains and
	//
	//   // The above uses a bunch of defaults that you should be aware of
	//   // AutoChain defaults to
	// 	//Build a chain
	// 	rc, err := cl.BuildChain(uri, "C*", us)
	// 	if err != nil {
	// 		bail(err)
	// 	}
	// 	fmt.Println("BC returned")
	// 	pac := ""
	// 	for chain := range rc {
	// 		fmt.Printf("Got chain: %v", chain.Hash)
	// 		pac = chain.Hash
	// 	}
	// 	fmt.Println("chain build rv")
	// 	if pac == "" {
	// 		bail("Could not find a PAC")
	// 	}
	// 	//Subscribe
	// 	msgchan, err := cl.Subscribe(&bw2bind.SubscribeParams{
	// 		URI:                uri,
	// 		PrimaryAccessChain: pac,
	// 		ElaboratePAC:       bw2bind.ElaborateFull,
	// 	})
	// 	if err != nil {
	// 		bail(err)
	// 	}
	//
	// 	fmt.Println("subscribe finished")
	// 	if err != nil {
	// 		bail(err)
	// 	}
	// 	for m := range msgchan {
	// 		m.Dump()
	// 	}
	// 	fmt.Println("channel ended")
}
