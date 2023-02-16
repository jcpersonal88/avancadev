FROM golang:1.20.1

WORKDIR /go/src/fullcycle

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN GOOS=linux go build c.go

EXPOSE 9092

ENTRYPOINT [ "./c" ]