FROM golang:1.19

RUN apt-get update

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

CMD ["/app/main"]
