FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /app

COPY go.* .

RUN go mod tidy

COPY . .

RUN go build -o stresser ./cmd/cli/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/stresser .

ENTRYPOINT [ "./stresser" ]