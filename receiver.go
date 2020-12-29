package main

import (
	"fmt"
	"github.com/TannerKvarfordt/Simple-Go-Networking/network_info"
	"github.com/TannerKvarfordt/Simple-Go-Networking/protobuf/protobuild/simple_msg"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"io"
	"net"
)

type MsgHandler = func(b []byte) error

type Receiver struct {
	socket      net.Conn
	msgHandlers map[string][]MsgHandler
}

func (receiver *Receiver) RegisterMsgHandler(name string, handler MsgHandler) {
	receiver.msgHandlers[name] = append(receiver.msgHandlers[name], handler)
}

func (receiver *Receiver) Receive(ch chan *any.Any) {
	for {
		serializedMsg := make([]byte, network_info.MAX_DATAGRAM_SIZE_BYTES)
		n, err := receiver.socket.Read(serializedMsg)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Socket closed by sender")
				close(ch)
				return
			}
			panic(err)
		}
		msg := &any.Any{}
		err = proto.Unmarshal(serializedMsg[:n], msg)
		if err != nil {
			panic(err)
		}
		fmt.Println("Received", msg.TypeUrl, "message")
		ch <- msg
	}
}

func handleSimpleMsg(serializedMsg []byte) error {
	msg := &simple_msg.SimpleMsg{}
	err := proto.Unmarshal(serializedMsg, msg)
	if err != nil {
		return err
	}

	fmt.Println("Handled SimpleMsg =", msg.String())
	return nil
}

func main() {
	fmt.Println("Starting receiver...")

	listener, err := net.Listen(network_info.PROTOCOL, network_info.SENDER_FULL)
	if err != nil {
		fmt.Println("Failed to establish listener with error:", err)
		return
	}

	fmt.Println("Waiting for a connection...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection with error:", err)
		return
	}
	fmt.Println("Connected")

	receiver := &Receiver{conn, make(map[string][]MsgHandler)}
	receiver.RegisterMsgHandler("SimpleMsg", handleSimpleMsg)

	ch := make(chan *any.Any)
	go receiver.Receive(ch)
	for msg := range ch {
		if handlers, ok := receiver.msgHandlers[msg.TypeUrl]; ok {
			// Known message type
			for _, handler := range handlers {
				err = handler(msg.Value)
				if err != nil {
					fmt.Println("Error handling msg of type", msg.TypeUrl, ":", err)
				}
			}
		} else {
			// Unhandled message type
			fmt.Println("Received unhandled message type", msg.TypeUrl, "-- ignoring")
		}
	}
}
