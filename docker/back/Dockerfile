FROM golang:1.21

RUN go version

WORKDIR /app

COPY ../../back/go.mod ../../back/go.sum  ./

RUN go mod download

CMD ["./back"]
EXPOSE 8080