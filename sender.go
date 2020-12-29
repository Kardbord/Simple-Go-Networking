package main

import (
	"fmt"
	"github.com/TannerKvarfordt/Simple-Go-Networking/network_info"
	"github.com/TannerKvarfordt/Simple-Go-Networking/protobuf/protobuild/simple_msg"
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"net"
	"time"
)

type Sender struct {
	socket net.Conn
}

func (sender *Sender) Send(msg descriptor.Message) error {
	_, desc := descriptor.ForMessage(msg)
	serializedMsg, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	generic := &any.Any{TypeUrl: *desc.Name, Value: serializedMsg}

	serializedGeneric, err := proto.Marshal(generic)
	if err != nil {
		return err
	}

	n, err := sender.socket.Write(serializedGeneric)
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

	msg := &simple_msg.SimpleMsg{SimpleString: "Test msg"}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		err = sender.Send(msg)
		if err != nil {
			fmt.Println("Failed to send msg with error:", err)
		}
	}
	_ = sender.socket.Close()
}
