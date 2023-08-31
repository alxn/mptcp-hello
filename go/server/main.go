// main for an example MP-TCP Server.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
)

func main() {
	lc := net.ListenConfig{}
	lc.SetMultipathTCP(true)
	ln, err := lc.Listen(context.Background(), "tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to start listener: %v", err)
	}
	//defer ln.Close()
	log.Printf("Listening on %v\n", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Accept error: %v\n", err)
		}
		log.Printf("Accepted from: %v\n", conn.RemoteAddr())
		go func() {
			//defer conn.Close()
			buf := make([]byte, 1024)
			for {
				_, err := conn.Read(buf)
				if err != nil {
					log.Printf("Read error: %v\n", err)
					return
				}
				_, err = conn.Write(buf)
				if err != nil {
					log.Printf("Write error: %v\n", err)
					return
				}
				fmt.Printf("!")
			}
		}()
	}
}
