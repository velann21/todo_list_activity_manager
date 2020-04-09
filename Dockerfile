FROM golang:latest AS builder
ADD . /app/backend
RUN ls /app/backend
WORKDIR /app/backend
RUN go mod download
RUN CGO_ENABLED=0 go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main /app/backend/cmd/activity_manager_bootstrap.go

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 8086
EXPOSE 2112


