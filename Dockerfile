FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o out

EXPOSE 8080

CMD ["go"]
