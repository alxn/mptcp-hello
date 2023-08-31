// main for an example MP-TCP Client.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"slices"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(99))
	d := net.Dialer{}
	d.SetMultipathTCP(true)
	conn, err := d.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatalf("Dial Error %v\n", err)
	}
	defer conn.Close()
	wbuf := make([]byte, 1024)
	rbuf := make([]byte, 1024)
	for {
		_, err := r.Read(wbuf)
		if err != nil {
			log.Fatalf("Random Read: %v\n", err)
		}
		_, err = conn.Write(wbuf)
		if err != nil {
			log.Fatalf("Write error: %v\n", err)
		}
		_, err = conn.Read(rbuf)
		if err != nil {
			log.Fatalf("Read error: %v\n", err)
		}
		if !slices.Equal(wbuf, rbuf) {
			log.Fatalf("Slices not equal!\n")
		}
		fmt.Print("!")
		time.Sleep(time.Second)
	}
}
