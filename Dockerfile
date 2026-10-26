FROM golang:1.26-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./api .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./api/cmd/api

FROM apline:3.22
WORKDIR /app
COPY --from=builder /api ./api
EXPOSE 8000
CMD ["./api"]
