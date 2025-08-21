package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	. "own-redis/internal"
)

func main() {
	flagport := flag.Int("port", 8080, "Port number")
	flaghelp := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *flaghelp {
		fmt.Println("Usage: own-redis [--port <N>]\nOptions:\n  --help\n  --port N\n  use nc 0.0.0.0 port")
		return
	}

	addr := fmt.Sprintf(":%d", *flagport)
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db := NewDB()
	buf := make([]byte, 1024)

	log.Printf("Server started on port %d", *flagport)

	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("Error reading: %v", err)
			continue
		}

		response := db.ProcessCommand(string(buf[:n]))
		conn.WriteTo([]byte(response), addr)
	}
}
