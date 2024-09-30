.PHONY: build clean

BINARY_NAME=bin/Project.exe
MAIN_PATH=cmd/Project/main.go

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

clean:
	rm -f $(BINARY_NAME)
