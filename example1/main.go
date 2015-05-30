package main

import (
	"fmt"
	"runtime"

	"github.com/immesys/bw2/objects"
	"github.com/immesys/bw2bind"
)

func main() {
	bw, err := bw2bind.Connect("127.0.0.1:28589")
	if err != nil {
		fmt.Println("Got error:", err)
		return
	}

	//Create an entity for ourselves
	eVK, keyblob, err := bw.CreateEntity(&bw2bind.CreateEntityParams{
		Comment: "for the win",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Entity created: ", eVK)

	//Set our entity
	err = bw.SetEntity(keyblob)
	if err != nil {
		fmt.Println("error: ", err.Error())
		panic(err)
	}

	//Publish to bunker
	bnkr := "uHn6zvvu46UdKgno7lpltIusjBnPcG_gMuFOPtnorUA="
	uri := bnkr + "/foo/bar"

	//Start subscriber
	msgs, err := bw.Subscribe(&bw2bind.SubscribeParams{
		URI:          uri,
		ElaboratePAC: "full",
		DoVerify:     true,
	})
	if err != nil {
		fmt.Println("error on subscribe: ", err)
		panic(err)
	}
	go func() {
		for sm := range msgs {
			fmt.Println("Got message From: ", sm.From)
			fmt.Println("             URI: ", sm.URI)
			fmt.Println("         Content: ", string(sm.POs[0].GetContent()))
		}
	}()

	//Publish a message
	err = bw.Publish(&bw2bind.PublishParams{
		URI:            uri,
		PayloadObjects: []objects.PayloadObject{bw2bind.MkPOString("hello world")},
		ElaboratePAC:   "full",
		DoVerify:       true,
	})
	if err != nil {
		panic(err)
	}

	for {
		runtime.Gosched()
	}
}
