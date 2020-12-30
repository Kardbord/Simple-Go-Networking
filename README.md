# Simple-Go-Networking
A simple learning experiment to try sending and receiving protocol buffers in Go.

# Build Instructions
## Prerequisites
* [Go](https://golang.org/)
* Google Protocol Buffers
  * [protoc](https://github.com/protocolbuffers/protobuf/tree/v3.14.0#protocol-compiler-installation)
    * You can build this from scratch or download a precompiled binary.
  * [protoc-gen-go](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)
* To use the build.py script, you'll also need a version of [python 3](https://www.python.org/downloads/).
## Building
Run the following commands:
```bash
go get github.com/TannerKvarfordt/Simple-Go-Networking
cd "${GOPATH}/src/github.com/TannerKvarfordt/Simple-Go-Networking"
./build.py
# To run the example:
cd build
./Simple-Go-Receiver &
./Simple-Go-Sender
```
