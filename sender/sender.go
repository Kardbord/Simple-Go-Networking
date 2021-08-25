package sender

import (
  "fmt"
  "net"

  "github.com/TannerKvarfordt/Simple-Go-Networking/network_info"
  "google.golang.org/protobuf/proto"
  "google.golang.org/protobuf/reflect/protoreflect"
  "google.golang.org/protobuf/types/known/anypb"
)

type sender struct {
  socket net.Conn
}

func NewSender(protocol string, toAddr string) (*sender, error) {
  conn, err := net.Dial(protocol, toAddr)
  if err != nil {
    return nil, fmt.Errorf("failed to establish %v connection to %v with error %v", network_info.PROTOCOL, network_info.RECEIVER_FULL, err)
  }
  s := sender{conn}
  return &s, nil
}

func (s *sender) Send(msg proto.Message) error {
  serializedMsg, err := proto.Marshal(msg)
  if err != nil {
    return err
  }

  msgName := protoreflect.Name(proto.MessageName(msg))
  generic := &anypb.Any{TypeUrl: string(msgName), Value: serializedMsg}

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
