package main

import (
	"fmt"
	"time"

	bw "github.com/immesys/bw2bind"
)

func main() {
	cl := bw.ConnectOrExit("")
	me := cl.SetEntityFileOrExit("ex2.key")

	// In the previous example, we used this to automatically use auto chain
	// meaning that the local router will attempt to find the permissions we
	// need for each command. In this example we will look at how we can build
	// chains manually and cache them, for increased performance or more
	// specificity (you may want a chain through a particular entity)
	// cl.OverrideAutoChainTo(true)

	// This says build chains that give Publish access to me on the
	// given URI. The result is a channel of chains
	resultchannel, err := cl.BuildChain("castle.bw2.io/bwtools/2", "P", me)

	// Lets count the number of chains that we got
	count := 0
	for _ = range resultchannel {
		count++
	}
	fmt.Println("We found", count, "chains")

	// Often, you really actually only want one chain. You can use
	// the utility method BuildAnyChain to get that. The OrExit variant
	// will print a neat error message and exit if it fails You will
	// notice the URI is for example 1. If you are running example 1
	// you can see these messages:
	chain := cl.BuildAnyChainOrExit("castle.bw2.io/bwtools/1", "P", me)
	po := bw.CreateStringPayloadObject("Hello from example 2")

	// Because we omitted the OverrideAutoChain call, AutoChain will
	// default to false. We must therefore include the hash of the chain
	// that we believe will give us permissions for the action. This is
	// much faster than autochain, and should be used by any real service
	// with the caveat that if the chain expires you should catch that
	// error and build another one whereas autochain will do that for you.
	// Future versions of BOSSWAVE will feature much faster chain building
	// based on block chain registries so this is only relevant for
	// 2.0.3 and earlier.
	err = cl.Publish(&bw.PublishParams{
		URI:                "castle.bw2.io/bwtools/1",
		PayloadObjects:     []bw.PayloadObject{po},
		PrimaryAccessChain: chain.Hash,
	})
	fmt.Println("Published with PAC, err was: ", err)

	// Without the PAC, you will encounter some errors. Here are examples:
	err = cl.Publish(&bw.PublishParams{
		URI:            "castle.bw2.io/bwtools/1",
		PayloadObjects: []bw.PayloadObject{po},
	})
	fmt.Println("Without PAC, error was: ", err)
	// You should see a message that says the local router attempted to
	// elaborate the PAC but you didn't tell it what the hash was. This
	// is because the default elaboration is full which means that the
	// messages the local router sends contain a fully self-standing proof
	// of permissions. We can instead tell the local router to not try
	// elaborating, and we will get a different error:
	err = cl.Publish(&bw.PublishParams{
		URI:            "castle.bw2.io/bwtools/1",
		PayloadObjects: []bw.PayloadObject{po},
		ElaboratePAC:   bw.ElaborateNone,
	})
	fmt.Println("Without elaboration, error was: ", err)
	// It should say that the message failed verification.

	// If for example we were to try sending a message with an invalid PAC
	// we would get an error like this:
	err = cl.Publish(&bw.PublishParams{
		URI:            "castle.bw2.io/bwtools/1",
		PayloadObjects: []bw.PayloadObject{po},
		ElaboratePAC:   bw.ElaborateNone,
		// Normally the local router verifies your outgoing messages. For the
		// sake of example, lets not do that and see what the remote router
		// says
		DoNotVerify:        true,
		PrimaryAccessChain: "This_Is_Clearly_An_Invalid_PAC_But_A_Hash__=",
	})
	fmt.Println("With an invalid PAC, error was: ", err)
	// It should say "could not resolve PAC". This message is coming from the
	// designated router (not local), telling you why it rejected your message.

	// Lets send some good messages to play with example 1
	for {
		err = cl.Publish(&bw.PublishParams{
			URI:                "castle.bw2.io/bwtools/1",
			PayloadObjects:     []bw.PayloadObject{po},
			PrimaryAccessChain: chain.Hash,
		})
		fmt.Println("Published with PAC, err was: ", err)
		time.Sleep(5 * time.Second)
	}

}
