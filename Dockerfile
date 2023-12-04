FROM ysicing/god AS builder

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod download

COPY . ./

RUN go build -v -o pangu .

FROM ysicing/debian

WORKDIR /app

COPY --from=builder /app/pangu /app/pangu

ENTRYPOINT ["/app/pangu"]
