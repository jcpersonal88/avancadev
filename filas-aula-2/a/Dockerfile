FROM golang:1.23.3

WORKDIR /go/src/fullcycle

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN GOOS=linux go build a.go

EXPOSE 9090

ENTRYPOINT [ "./a" ]