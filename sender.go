package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"protobuf-networking/network_info"
	"protobuf-networking/protobuf/protobuild/SimpleMsg"
)

type Sender struct {
	socket net.Conn
}

func (sender *Sender) Send() error {
	msg := &SimpleMsg.SimpleMsg{SimpleString: "Test msg"}
	serializedMsg, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("Failed to Marshal msg", msg)
	}
	_, err = sender.socket.Write(serializedMsg)
	return err
}

func main() {
	fmt.Println("Starting sender...")
	conn, err := net.Dial("tcp", network_info.RECEIVER_FULL)
	if err != nil {
		fmt.Println("Failed to establish TCP connection to", network_info.RECEIVER_FULL, "with error", err)
	}
	sender := &Sender{conn}
	err = sender.Send()
	if err != nil {
		fmt.Println("Failed to send msg with error:", err)
	}
}
