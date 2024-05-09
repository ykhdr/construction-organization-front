FROM golang:1.22.1-alpine as builder
LABEL authors="ykhdr"

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd cmd
COPY internal internal

RUN go build -o construction-ogranization-front ./cmd/main.go

FROM alpine:3.19.1
COPY --from=builder /build/construction-ogranization-front /bin/construction-ogranization-front

COPY templates /templates
COPY static /static

ENTRYPOINT ["/bin/construction-ogranization-front"]