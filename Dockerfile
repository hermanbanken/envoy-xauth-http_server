FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/server cmd/introspector/main.go

FROM scratch
WORKDIR /
COPY --from=builder /bin/server /bin/server
ENTRYPOINT ["/bin/server"]