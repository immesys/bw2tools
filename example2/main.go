package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"

	"github.com/immesys/bw2/internal/crypto"
)

func main() {
	roots := x509.NewCertPool()
	conn, err := tls.Dial("tcp", "127.0.0.1:4514", &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            roots,
	})
	fmt.Println("done")
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	cs := conn.ConnectionState()
	fmt.Printf("PeerCertificates: %+v\n", cs.PeerCertificates)
	for _, c := range cs.PeerCertificates {
		fmt.Println(c.Subject)
		fmt.Println(c.Signature)
	}
	proof := make([]byte, 96)
	_, err = io.ReadFull(conn, proof)
	if err != nil {
		panic("Failed to read proof;" + err.Error())
	}
	proofOK := crypto.VerifyBlob(proof[:32], proof[32:], cs.PeerCertificates[0].Signature)
	fmt.Println("Verifying with VK: ", proof[:32])
	fmt.Println("Signature: ", proof[32:])
	fmt.Println("Payload: ", cs.PeerCertificates[0].Signature)
	fmt.Println("Proof OK: ", proofOK)
	conn.Close()
}
