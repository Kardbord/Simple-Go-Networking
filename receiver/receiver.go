package receiver

import (
  "fmt"
  "io"
  "net"
  "sync"

  "github.com/TannerKvarfordt/Simple-Go-Networking/network_info"
  "google.golang.org/protobuf/proto"
  "google.golang.org/protobuf/types/known/anypb"
)

type MsgHandler = func(b []byte) error

type receiver struct {
  socket      net.Conn
  msgHandlers map[string][]MsgHandler
  startOnce   sync.Once
}

func NewReceiver(protocol string, fromAddr string) (*receiver, error) {
  listener, err := net.Listen(protocol, fromAddr)
  if err != nil {
    return nil, fmt.Errorf("failed to establish listener: %v", err)
  }

  fmt.Println("Waiting for a connection...")
  conn, err := listener.Accept()
  if err != nil {
    return nil, fmt.Errorf("failed to accept connection: %v", err)
  }
  fmt.Println("Connected")

  r := receiver{conn, make(map[string][]MsgHandler), sync.Once{}}
  return &r, nil
}

func (r *receiver) receiveRoutine(ch chan *anypb.Any) {
  for {
    serializedMsg := make([]byte, network_info.MAX_DATAGRAM_SIZE_BYTES)
    n, err := r.socket.Read(serializedMsg)
    if err != nil {
      if err == io.EOF {
        fmt.Println("Socket closed by sender")
        close(ch)
        return
      }
      panic(err)
    }
    msg := &anypb.Any{}
    err = proto.Unmarshal(serializedMsg[:n], msg)
    if err != nil {
      panic(err)
    }
    fmt.Println("Received", msg.TypeUrl, "message")
    ch <- msg
  }
}

func (r *receiver) handleRoutine(ch chan *anypb.Any) {
  for msg := range ch {
    if handlers, ok := r.msgHandlers[msg.TypeUrl]; ok {
      // Known message type
      for _, handler := range handlers {
        err := handler(msg.Value)
        if err != nil {
          fmt.Println("Error handling msg of type", msg.TypeUrl, ":", err)
        }
      }
    } else {
      // Unhandled message type
      fmt.Println("Received unhandled message type \"", msg.TypeUrl, "\" -- ignoring it")
    }
  }
}

func (r *receiver) StartReceiver(blockOnThisCall bool) {
  r.startOnce.Do(func() {
    ch := make(chan *anypb.Any)
    if blockOnThisCall {
      go r.receiveRoutine(ch)
      r.handleRoutine(ch)
    } else {
      go r.handleRoutine(ch)
      go r.receiveRoutine(ch)
    }
  })
}

func (r *receiver) RegisterMsgHandler(name string, handler MsgHandler) {
  r.msgHandlers[name] = append(r.msgHandlers[name], handler)
}

