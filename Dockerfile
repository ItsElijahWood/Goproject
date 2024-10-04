# Make sure to run the image with "docker run -itd --rm -p 80:80 --name go go"

FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY /cmd/Project/main.go ./cmd/Project/main.go

COPY /cmd/Project/static ./cmd/Project/static

RUN go build -o bin/Project.exe ./cmd/Project/main.go

CMD ["./bin/Project.exe"]
