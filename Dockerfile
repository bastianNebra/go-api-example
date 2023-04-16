FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

EXPOSE 9000

CMD [ "./main" ]