# Simple-Go-Networking
A simple learning experiment to try sending and receiving protocol buffers in Go.

# Build Instructions
Since this is a project written in Go, obviously you will need [Go](https://golang.org/) installed before attempting to build it.
This project also has some additional dependencies. 
The simplest way to retrieve them and build is to use `make`. It will attempt to retrieve and install the dependencies for you. You must have the `GOPATH` and `GOBIN` environment variables set in order for this to work.
```bash
# This will display a warning about the missing simple_msg package.
# That's fine because we are going to build it later.
go get -u github.com/TannerKvarfordt/Simple-Go-Networking
cd "${GOPATH}/src/github.com/TannerKvarfordt/Simple-Go-Networking"
make
```
If that does not work or you want to do everything the hard way, manually install everything in the [Dependencies](#dependencies) section. Once that is done, run the following commands.
```bash
# This will display a warning about the missing simple_msg package.
# That's fine because we are going to build it later.
go get -u github.com/TannerKvarfordt/Simple-Go-Networking
cd "${GOPATH}/src/github.com/TannerKvarfordt/Simple-Go-Networking"
./build.py -b
```
If you don't even want to make use of the `build.py` script, you can run the following commands instead.
```bash
# This will display a warning about the missing simple_msg package.
# That's fine because we are going to build it later.
go get -u github.com/TannerKvarfordt/Simple-Go-Networking
cd "${GOPATH}/src/github.com/TannerKvarfordt/Simple-Go-Networking"
protoc -I=protobuf --go_out=protobuf protobuf/SimpleMsg.proto
go build -v -o build/Simple-Go-Sender sender-main.go
go build -v -o ${GOBIN}/Simple-Go-Sender sender-main.go
go build -v -o build/Simple-Go-Receiver receiver-main.go
go build -v -o ${GOBIN}/Simple-Go-Receiver receiver-main.go
```
Once you have built everything, you can run the example with the commands below.
```bash
# If $GOBIN is in your PATH
Simple-Go-Receiver &
Simple-Go-Sender
# Otherwise
cd $GOBIN
./Simple-Go-Receiver &
./Simple-Go-Sender
```
# Dependencies
* [Go](https://golang.org/)
* Google Protocol Buffers
  * [protoc](https://github.com/protocolbuffers/protobuf/tree/v3.14.0#protocol-compiler-installation)
    * You can build this from scratch or download a precompiled binary.
  * [protoc-gen-go](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)
  
To use the `build.py` script, you will also need:
* [python 3](https://www.python.org/downloads/)
