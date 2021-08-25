SENDER_MAIN = sender-main.go
RECV_MAIN = receiver-main.go
BUILD_DIR = build
SENDER_TARGET = Simple-Go-Sender
RECV_TARGET = Simple-Go-Receiver
PROTO_SRC = protobuf
PROTO_BUILD_DIR = $(PROTO_SRC)/protobuild

PROTOC := $(shell which protoc)
# If protoc isn't on the path, set it to a target that's never up to date, so
# the install command always runs.
ifeq ($(PROTOC),)
	PROTOC = must-rebuild
endif

PROTOC_GEN_GO := $(GOBIN)/protoc-gen-go

# Figure out which machine we're running on.
UNAME := $(shell uname)

all: $(SENDER_TARGET) $(RECV_TARGET)

$(PROTOC):
	# Run the right installation command for the operating system.
ifeq ($(UNAME), Darwin)
	brew install protobuf
endif
ifeq ($(UNAME), Linux)
	sudo apt-get install protobuf-compiler
endif

# If $GOPATH/bin/protoc-gen-go does not exist, we'll run this command to install it.
$(PROTOC_GEN_GO):
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$(SENDER_TARGET): $(SENDER_MAIN) SimpleMsg.pb.go
	go build -v -o $(BUILD_DIR)/$(SENDER_TARGET) $(SENDER_MAIN)
	go build -v -o $(GOBIN)/$(SENDER_TARGET) $(SENDER_MAIN)

$(RECV_TARGET): $(RECV_MAIN) SimpleMsg.pb.go
	go build -v -o $(BUILD_DIR)/$(RECV_TARGET) $(RECV_MAIN)
	go build -v -o $(GOBIN)/$(RECV_TARGET) $(RECV_MAIN)

SimpleMsg.pb.go: $(PROTO_SRC)/SimpleMsg.proto | $(PROTOC_GEN_GO) $(PROTOC)
	protoc -I=$(PROTO_SRC) --go_out=$(PROTO_SRC) $(PROTO_SRC)/SimpleMsg.proto

clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(PROTO_BUILD_DIR)
	rm  -f $(GOBIN)/$(SENDER_TARGET)
	rm  -f $(GOBIN)/$(RECV_TARGET)
