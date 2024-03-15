FROM registry.apps.xplat.fis.com.vn/library/golang:1.18
WORKDIR /app
COPY . .
RUN env GOOS=linux GOARCH=amd64 go build /app/cmd/main/main.go
EXPOSE 8080

CMD ["./main"]