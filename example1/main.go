package main

import (
	"fmt"

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
	fmt.Println("Keyfile: ", keyblob)
	//Set our entity
	err = bw.SetEntity(keyblob)
	if err != nil {
		fmt.Println("error: ", err.Error())
		panic(err)
	}
	//Publish to our own namespace
	uri := eVK + "/foo/bar"
	err = bw.Publish(&bw2bind.PublishParams{
		URI:            uri,
		PayloadObjects: []objects.PayloadObject{bw2bind.MkPOString("hello world")},
		DoVerify : true,
	})
	if err != nil {
		panic(err)
	}
}
