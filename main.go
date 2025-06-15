package main

import (
	"fmt"
	"net"
)

// Ensures gofmt doesn't remove the "net" import in stage 1 (feel free to remove this!)
var _ = net.ListenUDP

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	header := [12]byte{
		0x04, 0xD2, // ID 1234
		0x80, 0x00, // QR = 1, OPCODE = 4, AA= 1, TC= 1, RD = 1, RA=1, Z = 3, RCODE = 4
		0x00, 0x01,//QD COUNT // no of questions in the question section
		0x00, 0x00,//AN COUNT
		0x00, 0x00,//NS COUNT
		0x00, 0x00,//AR COUNT
	}
	question :=[]byte("\x0ccodecrafters\x02io\x00\x01\x00\x01")

	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		response := append(header[:], question)

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
