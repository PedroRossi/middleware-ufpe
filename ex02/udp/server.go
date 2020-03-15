package main

import (
	"log"
	"net"
)

func catHandler(pc net.PacketConn, addr net.Addr, buf []byte) {
	buf[2] |= 0x80 // Set QR bit
	pc.WriteTo(buf, addr)
}

func main() {
	pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		go catHandler(pc, addr, buf[:n])
	}
}
