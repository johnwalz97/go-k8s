FROM golang:latest as go-build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server

FROM scratch

COPY --from=go-build /app/server /server
EXPOSE 8080
CMD ["/server"]
