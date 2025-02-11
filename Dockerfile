FROM golang:1.22.2-alpine

WORKDIR /

COPY go.mod ./
COPY cmd ./cmd

RUN go build -o /bin/app ./cmd/app/

EXPOSE 8080

CMD ["/bin/app"]