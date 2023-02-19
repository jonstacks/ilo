FROM golang:1.20-alpine as builder

WORKDIR /ilo
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install ./cmd/...

FROM alpine
COPY --from=builder /go/bin/ilo-sweep /usr/local/bin/
COPY --from=builder /go/bin/ilo-server /usr/local/bin/
CMD ["ilo-sweep"]
