FROM golang:bookworm

WORKDIR /go/src/fullcycle

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN GOOS=linux go build b.go

EXPOSE 9091

ENTRYPOINT [ "./b" ]