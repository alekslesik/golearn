FROM golang:1.20

WORKDIR /root/go/src/github.com/alekslesik/golearn

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

# EXPOSE 9999

CMD ["app"]