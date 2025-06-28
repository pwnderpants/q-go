GO_FILES := main.go
BIN := ./build/q-go

all: $(BIN)

$(BIN):
	go build -o $(BIN) $(GO_FILES)

clean:
	rm -f $(BIN)

install:
	cp $(BIN) /usr/local/bin/q-go
