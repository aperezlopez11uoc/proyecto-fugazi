FROM golang:1.26

WORKDIR /app

COPY app/go.mod ./
RUN go mod download

COPY app/ .

RUN go build -o server main.go

EXPOSE 8080

ENV PORT=8080

CMD ["./server"]
