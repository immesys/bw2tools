package main

import (
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
	us, err := cl.SetEntityFile("ex5.key")
	if err != nil {
		bail(err)
	}
	uri := "michael.bw2.io/ex/5/foo"
	pac, e := cl.BuildAnyChain(uri, "PC", us)
	if pac == nil {
		fmt.Println("Could not get permissions: ", e)
		os.Exit(1)
	}
	fmt.Println("using pac: ", pac.Hash)

	//Read existing value
	vals, err := cl.Query(&bw.QueryParams{
		URI:                uri,
		PrimaryAccessChain: pac.Hash,
		ElaboratePAC:       bw.ElaborateFull,
	})
	if err != nil {
		bail(err)
	}
	fmt.Println("existing vals:")
	for rv := range vals {
		rv.Dump()
	}
	fmt.Println("-")

	//Persist some stuff
	err = cl.Publish(&bw.PublishParams{
		URI:                uri,
		PrimaryAccessChain: pac.Hash,
		DoVerify:           true,
		ElaboratePAC:       bw.ElaborateFull,
		Persist:            true,
		PayloadObjects:     []bw.PayloadObject{bw.CreateStringPayloadObject("Hello world")},
	})
	if err != nil {
		fmt.Println("finished persisting err: ", err.Error())
	} else {
		fmt.Println("Finished persisting, no error")
	}

}
