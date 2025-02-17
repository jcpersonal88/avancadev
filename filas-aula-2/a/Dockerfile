FROM golang:1.24.0

WORKDIR /go/src/fullcycle

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN GOOS=linux go build a.go

EXPOSE 9090

ENTRYPOINT [ "./a" ]