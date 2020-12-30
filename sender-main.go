package main

import (
  "fmt"
  "github.com/TannerKvarfordt/Simple-Go-Networking/network_info"
  "github.com/TannerKvarfordt/Simple-Go-Networking/sender"
  "github.com/TannerKvarfordt/Simple-Go-Networking/protobuf/protobuild/simple_msg"
  "time"
)

func main() {
  fmt.Println("Starting sender...")
  s, err := sender.NewSender(network_info.PROTOCOL, network_info.RECEIVER_FULL)
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
  _ = s.CloseSocket()
}
