package sender

import (
  "fmt"
  "github.com/TannerKvarfordt/Simple-Go-Networking/network_info"
  "github.com/golang/protobuf/descriptor"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/ptypes/any"
  "net"
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

func (s *sender) CloseSocket() error {
  return s.socket.Close()
}
