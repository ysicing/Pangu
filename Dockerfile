FROM ysicing/god AS builder

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod download

RUN go install github.com/mitchellh/gox@latest \
    && go install github.com/go-task/task/v3/cmd/task@latest

COPY . ./

RUN task

FROM ysicing/debian

WORKDIR /app

COPY --from=builder /app/_output/pangu_linux_amd64 /app/pangu

CMD ["/app/pangu", "server"]
