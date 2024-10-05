FROM golang:1.23-alpine

# run image using: "docker run -itd --rm -p 80:80 --name go go"

WORKDIR /app

COPY go.mod ./

# Fetch dependencies
RUN go mod download

# copy local files to container
COPY /cmd/Project/main.go ./cmd/Project/main.go
COPY /cmd/Project/static ./cmd/Project/static
COPY /internal ./internal
COPY . .

RUN go build -o bin/Project.exe ./cmd/Project/main.go

CMD ["./bin/Project.exe"]
