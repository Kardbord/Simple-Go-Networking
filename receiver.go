package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"protobuf-networking/network_info"
	"protobuf-networking/protobuf/protobuild/simple_msg"
)

type Receiver struct {
	socket net.Conn
}

func (receiver *Receiver) Receive() error {
	serializedMsg := make([]byte, 512)
	n, err := receiver.socket.Read(serializedMsg)
	if err != nil {
		return err
	}
	fmt.Println("Read", n, "bytes")
	msg := &simple_msg.SimpleMsg{}
	err = proto.Unmarshal(serializedMsg[:n], msg)
	if err != nil {
		return err
	}
	fmt.Println("Received msg", msg)
	return nil
}

func main() {
	fmt.Println("Starting receiver...")
	
	listener, err := net.Listen(network_info.PROTOCOL, network_info.SENDER_FULL)
	if err != nil {
		fmt.Println("Failed to establish listener with error:", err)
		return
	}
	
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection with error:", err)
		return
	}
	
	receiver := &Receiver{conn}
	err = receiver.Receive()
	if err != nil {
		fmt.Println("Failed to receive with error:", err)
	}
}
