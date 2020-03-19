package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"protobuf-networking/network_info"
	"protobuf-networking/protobuf/protobuild/simple_msg"
)

type Sender struct {
	socket net.Conn
}

func (sender *Sender) Send() error {
	msg := &simple_msg.SimpleMsg{SimpleString: "Test msg"}
	serializedMsg, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	n, err := sender.socket.Write(serializedMsg)
	fmt.Println("Wrote", n, "bytes")
	return err
}

func main() {
	fmt.Println("Starting sender...")
	conn, err := net.Dial(network_info.PROTOCOL, network_info.RECEIVER_FULL)
	if err != nil {
		fmt.Println("Failed to establish", network_info.PROTOCOL, "connection to", network_info.RECEIVER_FULL, "with error", err)
		return
	}
	sender := &Sender{conn}
	err = sender.Send()
	if err != nil {
		fmt.Println("Failed to send msg with error:", err)
	}
}
