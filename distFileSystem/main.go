package main

import (
	"distFileSyetem/p2p"
	"fmt"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListnerAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("we goochiee ")
	select {}
}
