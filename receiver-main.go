package main

import (
	"fmt"
	"github.com/TannerKvarfordt/Simple-Go-Networking/receiver"
  "github.com/TannerKvarfordt/Simple-Go-Networking/network_info"
  "github.com/TannerKvarfordt/Simple-Go-Networking/protobuf/protobuild/simple_msg"
  "github.com/golang/protobuf/proto"
)

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

	r, err := receiver.NewReceiver(network_info.PROTOCOL, network_info.SENDER_FULL)
  if err != nil {
    panic(err)
  }
  if r == nil {
    panic(fmt.Errorf("Receiver object was not created"))
  }

	r.RegisterMsgHandler("SimpleMsg", handleSimpleMsg)

	r.StartReceiver(true)
}
