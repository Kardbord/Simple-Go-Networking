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

type sender struct {
	socket net.Conn
}

func NewSender(protocol string, toAddr string) (*sender, error) {
	conn, err := net.Dial(protocol, toAddr)
	if err != nil {
		return nil, fmt.Errorf("Failed to establish", network_info.PROTOCOL, "connection to", network_info.RECEIVER_FULL, "with error", err)
	}
  s := sender{conn}
  return &s, nil
}

func (s *sender) Send(msg descriptor.Message) error {
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

	n, err := s.socket.Write(serializedGeneric)
	fmt.Println("Wrote", n, "bytes")
	return err
}

func main() {
	fmt.Println("Starting sender...")
	s, err := NewSender(network_info.PROTOCOL, network_info.RECEIVER_FULL)
  if err != nil {
    panic(err)
  }
  if s == nil {
    panic(fmt.Errorf("Sender object was not created"))
  }

	msg := &simple_msg.SimpleMsg{SimpleString: "Test msg"}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		err = s.Send(msg)
		if err != nil {
			fmt.Println("Failed to send msg:", err)
		}
	}
	_ = s.socket.Close()
}
